import type { PlasmoCSConfig, PlasmoRender } from "plasmo"

import { scrape_freelancers_action } from "~constants"
import { Locator } from "~locator/locator"
import { UpworkFreelancerScraper } from "~scraper/upwork-freelancer-scraper"
import type {
  Earnings_Info,
  FreelancerJobHistory,
  ProcessFreelancerJobHistoryResult,
  Scraped_Freelancer_Data,
  ScrapeFreelancerResponse
} from "~types/freelancer"

export const config: PlasmoCSConfig = {
  matches: ["https://www.upwork.com/*/applicants/*/applicants*"]
}

chrome.runtime.onMessage.addListener(
  function (request, sender, sendResponse) {
    if (request.action === scrape_freelancers_action) {
      console.log("Scraping Freelancers...")
      const locator = new Locator(2000);
      const UprankFreelancerScraper = new UpworkFreelancerScraper(
        locator
      )
      UprankFreelancerScraper.scrape().then((scraped_freelancers) => {
        sendResponse(scraped_freelancers)    
      })
    }
    return true;
  }
)