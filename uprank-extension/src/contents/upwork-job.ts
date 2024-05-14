import { parse } from "path"
import cssText from "data-text:~/style.css"
import type { PlasmoCSConfig, PlasmoRender } from "plasmo"

import { scrape_job_action } from "~constants"
import type { Job, ScrapedJobData, UnstableScrapedJobData } from "~types/job"
import { extractJobId, getCurrentTabUrl } from "~utils/url-functions"

export const config: PlasmoCSConfig = {
  matches: ["https://www.upwork.com/*/applicants/*/job-details/*"]
}
const MISSING_FIELD: string = "not found";

//todo: NEED TO ADD A CHECK TO SEE IF ALL FIELDS WERE FILLED.
//if they were not filled out, we need a way to allow a user to fill them out manually.
chrome.runtime.onMessage.addListener(async (request, sender, sendResponse) => {
  if (request.action === scrape_job_action) {
    const jobId = request.jobId;
    let unstableJobData: UnstableScrapedJobData = await scrapeJob();
    unstableJobData.job.id = jobId;
    sendResponse({ job: unstableJobData.job, missingFields: unstableJobData.missingFields});
  }
  return true; // Ensure to return true for asynchronous response handling
});


const scrapeJob = async (): Promise<UnstableScrapedJobData> => {

  const title =
    document.querySelector("h4.mb-2x").textContent.trim() || MISSING_FIELD

  const descriptionElement = document.querySelector(
    'div[data-v-3bf30560][data-v-316735dc][data-test="Description"] p.text-body-sm'
  )
  const descriptionText = descriptionElement
    ? descriptionElement.textContent
    : MISSING_FIELD

  const location =
    document.querySelector("span.text-light-on-muted").textContent.trim() ||
    MISSING_FIELD

  const paymentTypeString =
    document
      .querySelector("div.description[data-v-36cea51f]")
      .textContent.trim() || MISSING_FIELD
  const hourly = paymentTypeString === "Hourly"
  const fixed = paymentTypeString === "Fixed-price"

  const priceElements = document.querySelectorAll(
    'div[data-v-8d6ae40e][data-v-36cea51f][data-test="BudgetAmount"] p.m-0 strong'
  )

  const prices = Array.from(priceElements).map((element) =>
    element.textContent.trim()
  )

  const experience_level =
    document
      .querySelector('div.air3-icon.md[data-cy="expertise"][data-v-36cea51f]')
      .textContent.trim() || MISSING_FIELD
  const skillElements = document.querySelectorAll(
    'div.skills-list span[data-test="Skill"] a'
  )

  // Map NodeList to an array of skill names
  const skills = Array.from(skillElements).map((element) =>
    element.textContent.trim()
  )

  const job: ScrapedJobData = {
    id: "unset",
    title: title,
    location: location,
    description: descriptionText,
    skills: skills,
    experience_level: experience_level,
    hourly: hourly,
    fixed: fixed,
    hourly_rate: hourly
      ? prices.map((price) => {
          return convertCurrencyToFloat(price)
        })
      : [],
    fixed_rate: fixed ? convertCurrencyToFloat(prices[0]) : 0
  }
  //todo: implement check for missingFields
  //also todo: implement a way to detect missing fields?
  return {
    job:job,
    missingFields: false
  }
}


function convertCurrencyToFloat(currencyStr: string): number | null {
  const numericStr = currencyStr.replace(/[^0-9.]/g, "")
  const result = parseFloat(numericStr)
  return isNaN(result) ? null : result
}


