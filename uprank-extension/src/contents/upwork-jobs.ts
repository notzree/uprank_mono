import cssText from "data-text:~/style.css"
import type { PlasmoCSConfig, PlasmoRender } from "plasmo"

import type { Freelancer } from "~types/freelancer"

export const config: PlasmoCSConfig = {
  matches: ["https://www.upwork.com/*/applicants/*/job-details"]
}

chrome.runtime.onMessage.addListener(function (request, sender, sendResponse) {
  if (request.action === "fetchJob") {
    console.log("Activation command received")
    // const freelancerMap = scrapeFreelancers()
    sendResponse({ data: "freelancerMap"})
  }
})
//TODO: Clean this file up, then implement scraping of job details

const scrapeFreelancers = async () => {
  // Send a message to content script
  let expectedFreelancerCount = null
  let freelancerMap = {}
  console.log("Scraping job...")
  try {
    expectedFreelancerCount = scrapeFreelancerCount()
    console.log(`Number of proposals: ${expectedFreelancerCount}`)
  } catch (error) {
    console.error(error)
  }

  const freelancer_data = filterLocalStorage()
  for (let i = 0; i < freelancer_data.length; i++) {
    try {
      const freelancer_url_cipher =
        freelancer_data[i].application.profile.ciphertext
      if (freelancer_url_cipher in freelancerMap) {
        continue
      } else {
        const freelancer: Freelancer = {
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
          earningsInfo: freelancer_data[i].application.profile.earningsInfo,
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

        freelancerMap[freelancer_url_cipher] = freelancer
      }
    } catch (e) {
      console.info(e + "occured at " + i + "th iteration")
    }
  }
  //make sure the expect matches the actual count
  if (Object.keys(freelancerMap).length != expectedFreelancerCount) {
    console.log("discrepancy in the number of proposals") //todo: implement manual adjustment
    return
  }
  console.log(freelancerMap)
  return freelancerMap
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
