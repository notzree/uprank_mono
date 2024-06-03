import type { PlasmoCSConfig, PlasmoRender } from "plasmo"

import { scrape_freelancers_action } from "~constants"
import type {
  Earnings_Info,
  Scraped_Freelancer_Data,
  ScrapeFreelancerResponse,
  FreelancerJobHistory, ProcessFreelancerJobHistoryResult,
} from "~types/freelancer"


export const config: PlasmoCSConfig = {
  matches: ["https://www.upwork.com/*/applicants/*/applicants*"]
}

chrome.runtime.onMessage.addListener(
  async function (request, sender, sendResponse) {
    if (request.action === scrape_freelancers_action) {
      const scraped_freelancers: ScrapeFreelancerResponse =
        await scrapeFreelancers()
      sendResponse({
        freelancers: scraped_freelancers.freelancers,
        missingFields: scraped_freelancers.missing_fields,
        missingFreelancers: scraped_freelancers.missing_freelancers
      })
    }
    return true
  }
)

const scrapeFreelancers =
  async (): Promise<ScrapeFreelancerResponse> => {
    console.log("Scraping Freelancers...")
    const freelancer_job_map = await preProcessFreelancers()
   

    let expectedFreelancerCount = null
    let unstableFreelancerMap = {}
    try {
      expectedFreelancerCount = scrapeFreelancerCount()
    } catch (error) {
      console.error(error)
    }

    const freelancer_data = filterLocalStorage()
    for (let i = 0; i < freelancer_data.length; i++) {
      const freelancer_url_cipher =
        freelancer_data[i].application.profile.ciphertext
      if (freelancer_url_cipher in unstableFreelancerMap) {
        continue
      } else {
        const scraped_freelancer_data: Scraped_Freelancer_Data = {
          name: freelancer_data[i].application.profile.shortName,
          location: {
            city: freelancer_data[i].application.profile.location.city,
            country: freelancer_data[i].application.profile.location.country,
            timezone: freelancer_data[i].application.profile.location.timezone
          },
          title: freelancer_data[i].application.profile.title,
          description: freelancer_data[i].application.profile.description,
          photo_url: freelancer_data[i].application.profile.photoUrl,
          hourly_charge_amount:
            freelancer_data[i].application.hourlyChargeRate.amount,
          hourly_charge_currency:
            freelancer_data[i].application.hourlyChargeRate.currencyCode,
          fixed_charge_amount:
            freelancer_data[i].application.fixedChargeAmount.amount,
          fixed_charge_currency:
            freelancer_data[i].application.fixedChargeAmount.currencyCode,
          earnings_info: flattenEarningsInfo(
            freelancer_data[i].application.profile.earningsInfo
          ),
          cv: freelancer_data[i].application.coverLetter,
          url: `https://www.upwork.com/freelancers/${freelancer_data[i].application.profile.ciphertext}`,
          ai_reccomended: freelancer_data[i].application.aiRecommended,
          invited: freelancer_data[i].application.invited,
          recent_hours: freelancer_data[i].application.profile.recentHours,
          total_hours: freelancer_data[i].application.profile.totalHours,
          total_portfolio_items:
            freelancer_data[i].application.profile.totalPortfolioItems,
          total_portfolio_v2_items:
            freelancer_data[i].application.profile.totalPortfolioV2Items,
          total_feedback: freelancer_data[i].application.profile.totalFeedback,
          recent_feedback:
            freelancer_data[i].application.profile.recentFeedback,
          top_rated_status:
            freelancer_data[i].application.profile.topRatedStatus ==
            "not_eligible"
              ? false
              : true,
          top_rated_plus_status:
            freelancer_data[i].application.profile.topRatedPlusStatus ==
            "not_eligible"
              ? false
              : true,
          sponsored: freelancer_data[i].application.sponsored,
          job_success_score:
            freelancer_data[i].application.profile.jobSuccessScore,
          reccomended: freelancer_data[i].application.recommended,
          skills: freelancer_data[i].application.profile.skills.map(
            (skill) => skill.prettyName
          ),
          attachements: freelancer_data[i].application.attachments.map(
            (attachment) => {
              return {
                name: attachment.name,
                link: `https://www.upwork.com${attachment.link}`
              }
            }
          ),
          work_history: null
        }

        unstableFreelancerMap[freelancer_url_cipher] = scraped_freelancer_data
      }
    }

    //make sure the expect matches the actual count
    const unstableFreelancerArray: Scraped_Freelancer_Data[] = Object.values(
      unstableFreelancerMap
    )
    for (const freelancer of unstableFreelancerArray) {
      freelancer.work_history = freelancer_job_map[freelancer.name]
    }

    console.log(unstableFreelancerArray);

    if (unstableFreelancerArray.length != expectedFreelancerCount) {
      console.log(
        `discrepancy in the number of proposals, expected ${expectedFreelancerCount} got ${unstableFreelancerArray.length}]`
      ) //todo: implement manual adjustment

      return {
        freelancers: unstableFreelancerArray,
        missing_fields: true,
        missing_freelancers:
          expectedFreelancerCount - unstableFreelancerArray.length
      }
    }


    return {
      freelancers: unstableFreelancerArray,
      missing_fields: false,
      missing_freelancers: 0
    }
  }

function filterLocalStorage() {
  // This regular expression matches 'modal' followed by digits, then '-expire-', followed by more digits
  const regex = /^modal\d+-expire-\d+$/

  const matchingValues = []

  // Access all keys in local storage
  for (let i = 0; i < localStorage.length; i++) {
    const key = localStorage.key(i)
    if (regex.test(key)) {
      // If the key matches the pattern, get the value from local storage
      const value = localStorage.getItem(key)
      matchingValues.push(JSON.parse(value))
    }
  }

  return matchingValues
}

async function preProcessFreelancers() {
  //this function serves 2 purposes
  //1. Clicks through all the freelancers to load their data into localstorage
  //2. Scrapes each freelancers job data and stores it in a hashnmap.
  await clickLoadMore() //ensure all freelancers are loaded onto the page
  const freelancer_job_map = {}
  const clickableDivs = await locate(
    querySelectorAll,
    'div[data-ev-tab="applicants"][data-test="ProposalTile"]'
  )
  for (const div of clickableDivs as any) {
    div.click()
    const popupSelector =
      ".air3-card.air3-card-sections.air3-card-outline.profile-outer-card.mb-4x"
    const closeButtonSelector =
      "button.m-0.p-0.air3-btn.air3-btn-link.d-none.d-md-block"
    const freeelancer_job_map_entry = await processFreelancerJobHistory(popupSelector, closeButtonSelector)
    if (freeelancer_job_map_entry) {
      freelancer_job_map[freeelancer_job_map_entry.name] = freeelancer_job_map_entry.jobs
    }
    await waitForClose(closeButtonSelector)
  }
  return freelancer_job_map
}
//resolves once a given selector is no longer present on the page

async function processFreelancerJobHistory(popupSelector, closeButtonSelector): Promise<ProcessFreelancerJobHistoryResult | null>{
  //this function has weird and unexpected behaviour. I need it to open the popup, then do some operations on that popup, then close the popup. Instead it instantly opens and closes it. It doesn't load any of the information on it.
  try {
    // Wait for the popup and close button to appear
    const freelancer_name_element = await locate(querySelector, 'h2[itemprop="name"]');
    if (freelancer_name_element === null) {
      console.log("no freelancer name found")
      return null
    }
    const freelancer_name = freelancer_name_element.textContent.trim();
    const popup = await locate(querySelector, popupSelector)
    const closeButton: HTMLButtonElement = await locate(
      querySelector,
      closeButtonSelector
    )
    if (popup && closeButton) {
        const workHistoryDivs = await locate(
          querySelectorAll,
          'div[id="jobs_completed_desktop"] div.assignments-item'
        )
        if (workHistoryDivs === null) {
          console.log("no work history found")
          closeButton.click()
          return null
        }
        const jobs = []
        for (const div of workHistoryDivs as any) {
          const linkElement: HTMLAnchorElement = await locate(
            nestedSelector,      
            "a.up-n-link",
            div,
            7000
          )
          const jobTitle = linkElement.textContent.trim()
          const dateElement = (
            await locate(querySelector, 'div[data-test="assignment-date"]')
          ).textContent.trim()

          linkElement.click();
          const jobPopupModal = await locate(querySelector, 'div.air3-modal-body', null, 7000);
          const jobCloseButton : HTMLButtonElement = await locate(querySelector, 'button.air3-modal-close[data-test="UpCModalClose"]');
          const jobDetailDiv = await locate(nestedSelector, 'div[data-test="job-details"]', jobPopupModal, 5000);

          if (jobDetailDiv === null) {
            const job = {
              title: jobTitle,
              date: dateElement,
              description: "No job description found",
              budget: "No budget found"
            }
            jobs.push(job)
            jobCloseButton.click();
            continue
          }

          const jobDescriptionElement = await locate(nestedSelector, 'div[data-ev-sublocation="!line_clamp"]', jobPopupModal);
          const jobDescription = jobDescriptionElement ? jobDescriptionElement.getAttribute('data-test-key') : "No job description found";

          const jobBudgetElement = await locate(nestedSelector, 'div[data-test="job-details"] strong', jobPopupModal);
          const jobBudget = jobBudgetElement ? jobBudgetElement.textContent.trim() : "No budget found";
          
          const totalEarnedElement = await locate(nestedSelector, 'div[data-test="assignment-summary"] strong', jobPopupModal);
          const totalEarned = totalEarnedElement ? parseFloat(totalEarnedElement.textContent.trim().replace("$", "")) : -1;


          const clientTotalSpendElement = await locate(nestedSelector, 'div.mt-6x strong.d-block', jobPopupModal);
          const clientTotalSpend = clientTotalSpendElement ? clientTotalSpendElement.textContent.trim() : "No client total spend found";

          const clientTotalHiresElement = await locate(nestedSelector, 'div.mt-6x small.text-light-on-inverse', jobPopupModal);
          const clientTotalHires = clientTotalHiresElement ? clientTotalHiresElement.textContent.trim() : "No client total hires found";

          const clientFeedbackElement = await locate(nestedSelector, 'div[data-test="assignment-client-feedback"] em', jobPopupModal);
          const clientFeedback = clientFeedbackElement ? clientFeedbackElement.textContent.trim() : "No client feedback found";

          const clientRatingElement = await locate(nestedSelector, 'div[data-ev-sublocation="!rating"] strong', jobPopupModal, 1000);
          const clientRating = clientRatingElement ? parseFloat(clientRatingElement.textContent.trim()) : -1;
          
          jobCloseButton.click();
          await waitForClose('button.air3-modal-close[data-test="UpCModalClose"]');

          const job: FreelancerJobHistory = {
            title: jobTitle,
            date: dateElement,
            description: jobDescription,
            budget: jobBudget,
            total_earned: totalEarned,
            client_total_spend: clientTotalSpend,
            client_total_hires: clientTotalHires,
            client_rating: clientRating,
            client_feedback: clientFeedback
          }
          jobs.push(job)
        }


        closeButton.click()

        return {
          name: freelancer_name,
          jobs: jobs
        }
    }
  } catch (error) {
    console.error("An error occurred:", error)
    //todo: implement better error handling for this.
    return null
  }
}

const flattenEarningsInfo = (earningsInfo): Earnings_Info => {
  const flattened = {}
  for (const key in earningsInfo) {
    flattened[key] = earningsInfo[key].value
  }
  return flattened as Earnings_Info
}

function clickLoadMore() {
  return new Promise<void>((resolve, reject) => {
    function attemptClick() {
      // Get the wrapping div
      const loadMoreDiv = document.querySelector(
        "div.text-center.py-4x.border-top"
      )

      // Check if the div exists
      if (loadMoreDiv) {
        const loadMoreButton: HTMLElement = loadMoreDiv.querySelector(
          "button.air3-btn.air3-btn-secondary"
        )

        if (loadMoreButton) {
          loadMoreButton.click()
          setTimeout(attemptClick, 1000) // Continue attempting to click
        } else {
          setTimeout(attemptClick, 1000) // Recheck for button in the same div after delay
        }
      } else if (!loadMoreDiv) {
        console.log("All freelancers loaded")
        resolve() // Resolve the promise when the specific div is not found
      }
    }

    attemptClick() // Start the clicking process
  })
}

function scrapeFreelancerCount() {
  // Select all span elements that could potentially contain the proposal count
  const spans = document.querySelectorAll("span.air3-tab-btn-text")

  // Initialize proposalCount to null in case no matching element is found
  let proposalCount = null

  // Iterate over the span elements to find the one that contains proposal information
  spans.forEach((span) => {
    if (span.textContent.includes("Proposals")) {
      // Extract the number from the text content of the span element
      const matches = span.textContent.match(/\d+/) // This regex matches one or more digits
      if (matches) {
        proposalCount = parseInt(matches[0], 10) // Convert the first matched number to an integer
      }
    }
  })

  // Return the proposal count or throw an error if it wasn't found
  if (proposalCount !== null) {
    return proposalCount
  } else {
    throw new Error("Proposal count not found.")
  }
}

const querySelector = (input) => document.querySelector(input)

const querySelectorAll = (input) => document.querySelectorAll(input)

const nestedSelector = (input, parent) => parent.querySelector(input)

async function locate<T>(
  query: (input: string, parent?: Element) => T,
  selector: string,
  parent?: Element,
  timeout: number = 2000
): Promise<T | null> {
  return new Promise((resolve, reject) => {
    const intervalId = setInterval(() => {
      const element = parent ? query(selector, parent) : query(selector);
      if (element) {
        if (
          (element instanceof NodeList && element.length !== 0) ||
          element instanceof Element
        ) {
          clearInterval(intervalId);
          clearTimeout(timeoutId);
          resolve(element);
        }
      }
    }, 100);

    const timeoutId = setTimeout(() => {
      clearInterval(intervalId);
      console.log(
        `Element with selector "${selector}" not found within ${timeout}ms`
      );
      resolve(null);
    }, timeout);
  });
}




async function waitForClose(selector: string): Promise<void> {
  return new Promise<void>((resolve) => {
    const intervalId = setInterval(() => {
      const closeButton = document.querySelector(selector)
      if (!closeButton) {
        clearInterval(intervalId)
        resolve()
      }
    }, 100)
  })
}
