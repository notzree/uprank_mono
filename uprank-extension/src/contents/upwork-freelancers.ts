import type { PlasmoCSConfig, PlasmoRender } from "plasmo"
import { scrape_freelancers_action } from "~constants";
import type { Earnings_Info, Scraped_Freelancer_Data, Unstable_Scraped_Freelancer_Data } from "~types/freelancer"

export const config: PlasmoCSConfig = {
  matches: ["https://www.upwork.com/*/applicants/*/applicants*"]
}


chrome.runtime.onMessage.addListener(async function(request, sender, sendResponse) {
  if (request.action === scrape_freelancers_action) {
      const unstableFreelancers: Unstable_Scraped_Freelancer_Data = await scrapeFreelancers();
      sendResponse({ freelancers: unstableFreelancers.freelancers, missingFields: unstableFreelancers.missing_fields, missingFreelancers: unstableFreelancers.missing_freelancers});
  }
  return true;
});

const scrapeFreelancers = async (): Promise<Unstable_Scraped_Freelancer_Data> => {
  console.log("Scraping Freelancers...")
  // await loadLocalStorageData();

  let expectedFreelancerCount = null
  let unstableFreelancerMap = {}
  try {
    expectedFreelancerCount = scrapeFreelancerCount()
  } catch (error) {
    console.error(error)
  }
  
  const freelancer_data = filterLocalStorage()
  for (let i = 0; i < freelancer_data.length; i++) {
    try {
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
          hourly_charge_amount: freelancer_data[i].application.hourlyChargeRate.amount,
          hourly_charge_currency: freelancer_data[i].application.hourlyChargeRate.currencyCode,
          fixed_charge_amount: freelancer_data[i].application.fixedChargeAmount.amount,
          fixed_charge_currency: freelancer_data[i].application.fixedChargeAmount.currencyCode,
          earnings_info: flattenEarningsInfo(freelancer_data[i].application.earningsInfo),
          cv: freelancer_data[i].application.coverLetter,
          url: `https://www.upwork.com/freelancers/${freelancer_data[i].application.profile.ciphertext}`,
          ai_reccomended: freelancer_data[i].application.aiRecommended,
          invited: freelancer_data[i].application.invited,
          recent_hours: freelancer_data[i].application.profile.recentHours,
          total_hours: freelancer_data[i].application.profile.totalHours,
          total_portfolio_items: freelancer_data[i].application.profile.totalPortfolioItems,
          total_portfolio_v2_items: freelancer_data[i].application.profile.totalPortfolioV2Items,
          total_feedback: freelancer_data[i].application.profile.totalFeedback,
          recent_feedback: freelancer_data[i].application.profile.recentFeedback,
          top_rated_status: freelancer_data[i].application.profile.topRatedStatus == "not_eligible" ? false : true,
          top_rated_plus_status: freelancer_data[i].application.profile.topRatedPlusStatus == "not_eligible" ? false : true,
          sponsored: freelancer_data[i].application.sponsored,
          job_success_score: freelancer_data[i].application.profile.jobSuccessScore,
          reccomended: freelancer_data[i].application.recommended,
          skills: freelancer_data[i].application.profile.skills.map((skill) => skill.prettyName),
          attachements: freelancer_data[i].application.attachments.map((attachment) => {
              return {
                  name: attachment.name,
                  link: `https://www.upwork.com${attachment.link}`
              };
          })
      };
      
        unstableFreelancerMap[freelancer_url_cipher] = scraped_freelancer_data
      }
    } catch (e) {
      console.info(e + "occured at " + i + "th iteration")
    }
  }
  //make sure the expect matches the actual count
  const unstableFreelancerArray: Scraped_Freelancer_Data[] = Object.values(unstableFreelancerMap)

  if (unstableFreelancerArray.length != expectedFreelancerCount) {
    console.log(`discrepancy in the number of proposals, expected ${expectedFreelancerCount} got ${unstableFreelancerArray.length}]`) //todo: implement manual adjustment

    return {
      freelancers: unstableFreelancerArray,
      missing_fields: true,
      missing_freelancers: expectedFreelancerCount - unstableFreelancerArray.length
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

async function loadLocalStorageData() {
  await clickLoadMore();
  
  const clickableDivs = document.querySelectorAll('div[data-ev-tab="applicants"][data-v-9cb7f262]');

  for (const div  of clickableDivs as any) {
    div.click();

    // Wait for the popup to appear
    await waitForPopupAndClickClose('.air3-card.air3-card-sections.air3-card-outline.profile-outer-card.mb-4x', 'button.m-0.p-0.air3-btn.air3-btn-link.d-none.d-md-block');


    // Wait a bit for the popup to close and system to settle before the next click
    await new Promise(resolve => setTimeout(resolve, 1000));
  }
}

function waitForPopupAndClickClose(popupSelector, closeButtonSelector) {
  return new Promise<void>(resolve => {
    const intervalId = setInterval(() => {
      const popup = document.querySelector(popupSelector);
      const closeButton = document.querySelector(closeButtonSelector);
      console.log(popup, closeButton);
      if (popup && closeButton) {
        closeButton.click();
        clearInterval(intervalId);
        resolve();
      }
    }, 100); // Check every 100 milliseconds
  });
}

const flattenEarningsInfo = (earningsInfo): Earnings_Info => {
  const flattened: Earnings_Info = {
    average_recent_earnings: 0,
    combined_average_recent_earnings: 0,
    combined_recent_earnings: 0,
    combined_total_earnings: 0,
    combined_total_revenue: 0,
    recent_earnings: 0,
    total_revenue: 0
};

  for (const key in earningsInfo) {
      if (earningsInfo.hasOwnProperty(key) && earningsInfo[key].hasOwnProperty('value')) {
          flattened[key] = earningsInfo[key].value;
      }
  }
  return flattened;
};


function clickLoadMore() {
  return new Promise<void>((resolve, reject) => {
    function attemptClick() {
      // Get the wrapping div
      const loadMoreDiv = document.querySelector('div.text-center.py-4x.border-top');

      // Check if the div exists
      if (loadMoreDiv) {
        const loadMoreButton: HTMLElement = loadMoreDiv.querySelector('button.air3-btn.air3-btn-secondary');

        if (loadMoreButton) {
          loadMoreButton.click();
          setTimeout(attemptClick, 1000); // Continue attempting to click
        } else {
          setTimeout(attemptClick, 1000); // Recheck for button in the same div after delay
        }
      } else if (!loadMoreDiv) {
        console.log("All freelancers loaded");
        resolve(); // Resolve the promise when the specific div is not found
      }
    }

    attemptClick(); // Start the clicking process
  });
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


