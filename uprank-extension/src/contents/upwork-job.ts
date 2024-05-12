import cssText from "data-text:~/style.css"
import type { PlasmoCSConfig, PlasmoRender } from "plasmo"

import { scrape_job_action } from "~constants"
import type { Freelancer } from "~types/freelancer"
import type { Job } from "~types/job"
import { extractJobId, getCurrentTabUrl } from "~utils/url-functions"

export const config: PlasmoCSConfig = {
  matches: ["https://www.upwork.com/*/applicants/*/job-details"]
}

chrome.runtime.onMessage.addListener(function (request, sender, sendResponse) {
  if (request.action === scrape_job_action) {
    console.log("Activation command received")
    const jobId = request.jobId //ID is passed from the popup.
    const jobData = scrapeJob() //excluding ID as content script can't access tabs api
    sendResponse({ data: "freelancerMap" })
  }
})
//TODO: Clean this file up, then implement scraping of job details

const scrapeJob = async () => {
  //scrape data
  // const url: string = await getCurrentTabUrl();
  // const jobId = extractJobId(url);

  const title_element = document.querySelector("h4.mb-2x")
  const title = title_element ? title_element.textContent : null

  const descriptionElement = document.querySelector(
    'div[data-v-3bf30560][data-v-316735dc][data-test="Description"] p.text-body-sm'
  )
  const descriptionText = descriptionElement
    ? descriptionElement.textContent
    : "Description not found"

  const locationElement = document.querySelector("span.text-light-on-muted")
  const location = locationElement ? locationElement.textContent : null

  const paymentTypeString =
    document
      .querySelector("div.description[data-v-36cea51f]")
      .textContent.trim() || null
  const hourly = paymentTypeString === "Hourly"
  const fixed = paymentTypeString === "Fixed-price"



  const priceElements = document.querySelectorAll(
    'div[data-v-8d6ae40e][data-v-36cea51f][data-test="BudgetAmount"] p.m-0 strong'
  )

  // Convert NodeList to an array and map to get trimmed text content
  const prices = Array.from(priceElements).map((element) =>
    element.textContent.trim()
  )

  console.log(descriptionText)
  console.log(title)
}
