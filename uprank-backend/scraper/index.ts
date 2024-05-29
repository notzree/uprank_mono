// For more information, see https://crawlee.dev/

import { QueueScrapeFreelancersReqest, SQSEvent } from './types.js';
import aws_chromium from '@sparticuz/chromium';
import { Configuration, PlaywrightCrawler } from 'crawlee';



export const handler = async (event: SQSEvent, _: any) => {
    
    const crawler = new PlaywrightCrawler({
        //@ts-ignore
        requestHandler: async ({ page, request, _ }) => {
            console.log(`Processing: ${request.url}`);
            const freelancerName = await page.locator('h2[itemprop="name"]').textContent();
            console.log(`Freelancer Name: ${freelancerName}`);
        },
    
        // Let's limit our crawls to make our tests shorter and safer.
        maxRequestsPerCrawl: 50,
        launchContext: {
            launchOptions: {
                 executablePath: await aws_chromium.executablePath(),
                 args: aws_chromium.args,
                 headless: true
            }
        }
    },
    new Configuration({
        persistStorage: false,
    }));
    
    const messages = event.Records.map((record) => JSON.parse(record.body));
   
    for (const message of messages){
        const request_data: QueueScrapeFreelancersReqest = JSON.parse(message.body);
        let urls = request_data.freelancers.map((freelancer) => freelancer.url);
        await crawler.run(urls);
    }

    return {
        statusCode: 200,
        body: "ok",
    }
}