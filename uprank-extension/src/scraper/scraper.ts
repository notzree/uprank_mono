import type { ScrapeFreelancerResponse } from "~types/freelancer";
// Scraper is a class that scrapes some data and returns it as an array of objects
interface Scraper<T> {
    scrape(): Promise<T>;
}

class UpworkFreelancerScraper implements Scraper<ScrapeFreelancerResponse> {
    //do stuff here
    async scrape(): Promise<ScrapeFreelancerResponse> {
        return { freelancers: [], missing_fields: false, missing_freelancers: 0 }
    }
}


