import type { PlasmoCSConfig, PlasmoRender } from "plasmo"
import { scrape_freelancers_action } from "~constants";
import type { EarningsInfo, ScrapedFreelancerData, UnstableScrapedFreelancerData } from "~types/freelancer"

export const config: PlasmoCSConfig = {
  matches: ["https://www.upwork.com/*/applicants/*/applicants*"]
}


chrome.runtime.onMessage.addListener(async function(request, sender, sendResponse) {
  if (request.action === scrape_freelancers_action) {
      const unstableFreelancers: UnstableScrapedFreelancerData = await scrapeFreelancers();
      sendResponse({ freelancers: unstableFreelancers.freelancers, missingFields: unstableFreelancers.missingFields, missingFreelancers: unstableFreelancers.missingFreelancers});
  }
  return true;
});

const scrapeFreelancers = async (): Promise<UnstableScrapedFreelancerData> => {
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
        const freelancer: ScrapedFreelancerData = {
          name: freelancer_data[i].application.profile.shortName,
          location: {
            city: freelancer_data[i].application.profile.location.city,
            country: freelancer_data[i].application.profile.location.country,
            timezone: freelancer_data[i].application.profile.location.timezone
          },
          title: freelancer_data[i].application.profile.title,
          description: freelancer_data[i].application.profile.description,
          photoUrl: freelancer_data[i].application.profile.photoUrl,
          hourlyChargeAmount:
            freelancer_data[i].application.hourlyChargeRate.amount,
          hourlyChargeCurrency:
            freelancer_data[i].application.hourlyChargeRate.currencyCode,
          fixedChargeAmount:
            freelancer_data[i].application.fixedChargeAmount.amount,
          fixedChargeCurrency:
            freelancer_data[i].application.fixedChargeAmount.currencyCode,
          earningsInfo: flattenEarningsInfo(freelancer_data[i].application.earningsInfo),
          cv: freelancer_data[i].application.coverLetter,
          url: `https://www.upwork.com/freelancers/${freelancer_data[i].application.profile.ciphertext}`,
          aiReccomended: freelancer_data[i].application.aiRecommended,
          invited: freelancer_data[i].application.invited,
          recentHours: freelancer_data[i].application.profile.recentHours,
          totalHours: freelancer_data[i].application.profile.totalHours,
          totalPortfolioItems:
            freelancer_data[i].application.profile.totalPortfolioItems,
          totalPortfolioV2Items:
            freelancer_data[i].application.profile.totalPortfolioV2Items,
          totalFeedback: freelancer_data[i].application.profile.totalFeedback,
          recentFeedback: freelancer_data[i].application.profile.recentFeedback,
          topRatedStatus:
            freelancer_data[i].application.profile.topRatedStatus ==
            "not_eligible"
              ? false
              : true,
          topRatedPlusStatus:
            freelancer_data[i].application.profile.topRatedPlusStatus ==
            "not_eligible"
              ? false
              : true,
          sponsored: freelancer_data[i].application.sponsored,
          jobSuccessScore:
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
          )
        }

        unstableFreelancerMap[freelancer_url_cipher] = freelancer
      }
    } catch (e) {
      console.info(e + "occured at " + i + "th iteration")
    }
  }
  //make sure the expect matches the actual count
  const unstableFreelancerArray: ScrapedFreelancerData[] = Object.values(unstableFreelancerMap)

  if (unstableFreelancerArray.length != expectedFreelancerCount) {
    console.log(`discrepancy in the number of proposals, expected ${expectedFreelancerCount} got ${unstableFreelancerArray.length}]`) //todo: implement manual adjustment

    return {
      freelancers: unstableFreelancerArray,
      missingFields: true,
      missingFreelancers: expectedFreelancerCount - unstableFreelancerArray.length
    }
  }
  return {
    freelancers: unstableFreelancerArray,
    missingFields: false,
    missingFreelancers: 0
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

const flattenEarningsInfo = (earningsInfo): EarningsInfo => {
  const flattened: EarningsInfo = {
      averageRecentEarnings: 0,
      combinedAverageRecentEarnings: 0,
      combinedRecentEarnings: 0,
      combinedTotalEarnings: 0,
      combinedTotalRevenue: 0,
      recentEarnings: 0,
      totalRevenue: 0
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


