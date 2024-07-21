import { Locator } from "~locator/locator";
import type {
    Earnings_Info,
    FreelancerJobHistory,
    ProcessFreelancerJobHistoryResult,
    Scraped_Freelancer_Data,
    ScrapeFreelancerResponse,
} from "~types/freelancer";

//Bit weird bc all of the private functions are intended to be ran on the browser, they don't return anyth and just change the state of the dom
export class UpworkFreelancerScraper
    implements Scraper<ScrapeFreelancerResponse>
{
    private locator: Locator;

    constructor(locator: Locator) {
        this.locator = locator;

        if (typeof document === "undefined") {
            throw new Error("scraper must be used in a browser environment");
        }
    }

    async scrape(): Promise<ScrapeFreelancerResponse> {
        let expectedFreelancerCount = null;
        let unstableFreelancerMap = {};
        try {
            expectedFreelancerCount = await this.scrapeFreelancerCount();
        } catch (error) {
            console.error(error);
        }
        const freelancer_job_map = await this.preProcessFreelancers();
        const freelancer_data = this.filterLocalStorage();
        for (let i = 0; i < freelancer_data.length; i++) {
            const freelancer_url_cipher =
                freelancer_data[i].application.profile.ciphertext;
            if (
                freelancer_url_cipher in unstableFreelancerMap ||
                !freelancer_job_map[
                    freelancer_data[i].application.profile.shortName
                ]
            ) {
                continue;
            } else {
                unstableFreelancerMap[freelancer_url_cipher] =
                    this.TransformFreelancerData(freelancer_data[i]);
            }
        }
        const unstableFreelancerArray: Scraped_Freelancer_Data[] =
            Object.values(unstableFreelancerMap);
        for (const freelancer of unstableFreelancerArray) {
            freelancer.work_history = freelancer_job_map[freelancer.name];
        }

        if (unstableFreelancerArray.length != expectedFreelancerCount) {
            console.log(
                `discrepancy in the number of proposals, expected ${expectedFreelancerCount} got ${unstableFreelancerArray.length}]`
            ); //todo: implement manual adjustment

            return {
                freelancers: unstableFreelancerArray,
                missing_fields: true,
                missing_freelancers:
                    expectedFreelancerCount - unstableFreelancerArray.length,
            };
        }

        return {
            freelancers: unstableFreelancerArray,
            missing_fields: false,
            missing_freelancers: 0,
        } as ScrapeFreelancerResponse;
    }

    private async preProcessFreelancers() {
        //this function serves 2 purposes
        //1. Clicks through all the freelancers to load their data into localstorage
        //2. Scrapes each freelancers job data and stores it in a hashnmap.
        await this.clickLoadMore(); //ensure all freelancers are loaded onto the page
        const freelancer_job_map = {};
        const clickableDivs = await this.locator.locateAll(
            'div[data-ev-tab="applicants"][data-test="ProposalTile"]'
        );
        for (const div of clickableDivs as any) {
            div.click();
            const popupSelector =
                ".air3-card.air3-card-sections.air3-card-outline.profile-outer-card.mb-4x";
            const closeButtonSelector =
                "button.m-0.p-0.air3-btn.air3-btn-link.d-none.d-md-block";
            const freelancer_job_map_entry =
                await this.processFreelancerJobHistory(
                    popupSelector,
                    closeButtonSelector
                );

            if (freelancer_job_map_entry) {
                freelancer_job_map[freelancer_job_map_entry.name] =
                    freelancer_job_map_entry.jobs;
            } else {
                freelancer_job_map[freelancer_job_map_entry.name] = [];
            }
            if (!freelancer_job_map_entry) {
                continue;
            }
            await this.locator.waitForClose(closeButtonSelector);
        }
        return freelancer_job_map;
    }

    private async processFreelancerJobHistory(
        popupSelector,
        closeButtonSelector
    ): Promise<ProcessFreelancerJobHistoryResult | null> {
        try {
            // Wait for the popup and close button to appear
            const freelancer_name_element = await this.locator.locate(
                "h2[itemprop='name']"
            );

            if (freelancer_name_element === null) {
                console.log("no freelancer name found");
                return {
                    name: "No name found",
                    jobs: [],
                };
            }
            const freelancer_name = freelancer_name_element.textContent.trim();
            const popup = await this.locator.locate(popupSelector);
            const closeButton = (await this.locator.locate(
                closeButtonSelector
            )) as HTMLButtonElement;
            if (popup && closeButton) {
                const jobs = [];
                let num_jobs_with_ratings = 0;
                let itr =0;
                const MAX_ITERATIONS = 25;
                while (num_jobs_with_ratings < 20 && itr < MAX_ITERATIONS) {
                    const workHistoryDivs = await this.locator.locateAll(
                        'div[id="jobs_completed_desktop"] div.assignments-item'
                    );
                    if (workHistoryDivs === null) {
                        console.log("no work history found");
                        closeButton.click();
                        return {
                            name: freelancer_name,
                            jobs: [],
                        };
                    }
                    for (const div of workHistoryDivs as any) {
                        const linkElement = (await this.locator.locate(
                            "a.up-n-link",
                            div
                        )) as HTMLAnchorElement;
                        const jobTitle = linkElement.textContent.trim();
                        linkElement.click();
                        const jobPopupModal = await this.locator.locate(
                            "div.air3-modal-body"
                        );
                        const jobCombinedDateElementSelector =
                            'div[data-test="assignment-date"]';
                        const jobCloseButtonSelector =
                            'button.air3-modal-close[data-test="UpCModalClose"]';
                        const jobDetailDivSelector =
                            'div[data-test="job-details"]';
                        const jobDescriptionElementSelector =
                            'div[data-ev-sublocation="!line_clamp"]';
                        const jobBudgetElementSelector =
                            'div[data-test="job-details"] div.air3-grid-container strong';
                        const totalEarnedElementSelector =
                            'div[data-test="assignment-summary"] div.mb-3x span';
                        const clientTotalSpendElementSelector =
                            'span[data-test="total-spent"] span';
                        const clientTotalHiresElementSelector =
                            "div.mt-6x small.text-light-on-inverse";
                        const clientFeedbackElementSelector =
                            'div[data-test="assignment-client-feedback"] em';
                        const clientRatingElementSelector =
                            'div[data-ev-sublocation="!rating"] span.sr-only';
                        const clientLocationElementSelector =
                            "div.span-md-5 div.mt-6x strong";
                        const result = await this.locator.locateMany<
                            Element | NodeList
                        >({
                            [jobCombinedDateElementSelector]: [
                                this.locator.querySelector,
                            ],
                            [jobCloseButtonSelector]: [
                                this.locator.querySelector,
                            ],
                            [jobDetailDivSelector]: [
                                this.locator.nestedSelector,
                                jobPopupModal,
                            ],
                            [jobDescriptionElementSelector]: [
                                this.locator.nestedSelector,
                                jobPopupModal,
                            ],
                            [jobBudgetElementSelector]: [
                                this.locator.nestedSelectorAll,
                                jobPopupModal,
                            ],
                            [totalEarnedElementSelector]: [
                                this.locator.nestedSelectorAll,
                                jobPopupModal,
                            ],
                            [clientTotalSpendElementSelector]: [
                                this.locator.nestedSelector,
                                jobPopupModal,
                            ],
                            [clientTotalHiresElementSelector]: [
                                this.locator.nestedSelector,
                                jobPopupModal,
                            ],
                            [clientFeedbackElementSelector]: [
                                this.locator.nestedSelector,
                                jobPopupModal,
                            ],
                            [clientRatingElementSelector]: [
                                this.locator.nestedSelector,
                                jobPopupModal,
                            ],
                            [clientLocationElementSelector]: [
                                this.locator.nestedSelector,
                                jobPopupModal,
                            ],
                        });
                        if (jobPopupModal === null) {
                            console.log("no job popup modal found");
                            closeButton.click();
                            return {
                                name: freelancer_name,
                                jobs: [],
                            };
                        }
                        const combined_dateElement = result[
                            jobCombinedDateElementSelector
                        ] as Element;
                        const combined_date =
                            combined_dateElement.textContent.trim();

                        const start_date = new Date(
                            combined_date.split(" - ")[0]
                        );
                        const end_date = new Date(
                            combined_date.split(" - ")[1]
                        );

                        const jobCloseButton = result[
                            jobCloseButtonSelector
                        ] as HTMLButtonElement;

                        const jobDetailDiv = result[jobDetailDivSelector];

                        if (jobDetailDiv === null) {
                            const job: FreelancerJobHistory = {
                                title: jobTitle,
                                start_date: start_date.toISOString(),
                                end_date: end_date.toISOString(),
                                description: null,
                                budget: null,
                                total_earned: null,
                                client_total_spend: null,
                                client_total_hires: null,
                                client_active_hires: null,
                                client_rating: null,
                                client_feedback: null,
                                client_location: null,
                            };
                            jobs.push(job);
                            jobCloseButton.click();
                            continue;
                        }

                        const jobDescriptionElement =
                            result[jobDescriptionElementSelector];
                        const jobDescription = jobDescriptionElement
                            ? (jobDescriptionElement as Element).getAttribute(
                                  "data-test-key"
                              )
                            : null;

                        const jobBudgetElementAll = result[
                            jobBudgetElementSelector
                        ] as NodeList;
                        var jobBudget = null;
                        for (const jobBudgetElement of jobBudgetElementAll) {
                            if (jobBudgetElement.textContent.includes("$")) {
                                jobBudget = parseFloat(
                                    jobBudgetElement.textContent
                                        .trim()
                                        .replace("$", "")
                                        .replace(",", "")
                                );
                                break;
                            }
                        }
                        // console.log(jobBudget);

                        const totalEarnedElementAll = result[
                            totalEarnedElementSelector
                        ] as NodeList;
                        
                        var totalEarned = null;
                        if (totalEarnedElementAll!=null){
                            for (const totalEarnedElement of totalEarnedElementAll) {
                                if (
                                    totalEarnedElement.textContent.includes("$") &&
                                    totalEarnedElement.textContent.includes(
                                        "earned"
                                    )
                                ) {
                                    totalEarned = parseFloat(
                                        totalEarnedElement.textContent
                                            .trim()
                                            .replace("$", "")
                                            .replace("earned", "")
                                            .replace(",", "")
                                    );
                                    break;
                                }
                            }
                        }
                            
                        const clientTotalSpendElement =
                            result[clientTotalSpendElementSelector];
                        const clientTotalSpend = clientTotalSpendElement
                            ? this.extractAmount(
                                  (
                                      clientTotalSpendElement as Element
                                  ).textContent.trim()
                              )
                            : null;

                        const clientTotalHiresElement =
                            result[clientTotalHiresElementSelector];
                        const clientTotalHiresString = clientTotalHiresElement
                            ? (
                                  clientTotalHiresElement as Element
                              ).textContent.trim()
                            : null;
                        const clientTotalHires = this.parseHireString(
                            clientTotalHiresString
                        ).total_hires;
                        const clientActiveHires = this.parseHireString(
                            clientTotalHiresString
                        ).active_hire;

                        const clientFeedbackElement =
                            result[clientFeedbackElementSelector];
                        const clientFeedback = clientFeedbackElement
                            ? (
                                  clientFeedbackElement as Element
                              ).textContent.trim()
                            : null;

                        const clientRatingElement =
                            result[clientRatingElementSelector];
                        const clientRating = clientRatingElement
                            ? parseFloat(
                                  (clientRatingElement as Element).textContent
                                      .trim()
                                      .split(" ")[2]
                              )
                            : null;

                        const clientLocationElement =
                            result[clientLocationElementSelector];
                        const clientLocation = clientLocationElement
                            ? (
                                  clientLocationElement as Element
                              ).textContent.trim()
                            : null;

                        jobCloseButton.click();
                        await this.locator.waitForClose(
                            'button.air3-modal-close[data-test="UpCModalClose"]'
                        );

                        const job: FreelancerJobHistory = {
                            title: jobTitle,
                            start_date: start_date.toISOString(),
                            end_date: end_date.toISOString(),
                            description: jobDescription,
                            budget: jobBudget,
                            total_earned: totalEarned,
                            client_total_spend: clientTotalSpend,
                            client_total_hires: clientTotalHires,
                            client_active_hires: clientActiveHires,
                            client_rating: clientRating,
                            client_feedback: clientFeedback,
                            client_location: clientLocation,
                        };
                        jobs.push(job);
                        num_jobs_with_ratings += 1;
                        if (num_jobs_with_ratings >= 50) {
                            closeButton.click();
                            return {
                                name: freelancer_name,
                                jobs: jobs,
                            };
                        }
                    }
                    const nextButtonElementSelector =
                        'div[id="jobs_completed_desktop"] button[data-ev-label="pagination_next_page"]';
                    const nextButton = (await this.locator.locate(
                        nextButtonElementSelector
                    )) as HTMLButtonElement;
                    if (!nextButton || nextButton.getAttribute("disabled") == "disabled") {
                        break;
                    }
                    else {
                        nextButton.click();
                        await new Promise((resolve) => setTimeout(resolve, 500));
                    }
                    itr++;
                }
                //Reached max jobs or no more jobs left
                closeButton.click();
                return {
                    name: freelancer_name,
                    jobs: jobs,
                };
            }
        } catch (error) {
            console.error("An error occurred:", error);
            //todo: implement better error handling for this.
            return {
                name: "error",
                jobs: [],
            };
        }
    }

    private clickLoadMore() {
        return new Promise<void>((resolve, reject) => {
            function attemptClick() {
                // Get the wrapping div
                const loadMoreDiv = document.querySelector(
                    "div.text-center.py-4x.border-top"
                );

                // Check if the div exists
                if (loadMoreDiv) {
                    const loadMoreButton: HTMLElement =
                        loadMoreDiv.querySelector(
                            "button.air3-btn.air3-btn-secondary"
                        );

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
    private TransformFreelancerData(
        freelancerData: any
    ): Scraped_Freelancer_Data {
        const scraped_freelancer_data: Scraped_Freelancer_Data = {
            name: freelancerData.application.profile.shortName,
            location: {
                city: freelancerData.application.profile.location.city,
                country: freelancerData.application.profile.location.country,
                timezone: freelancerData.application.profile.location.timezone,
            },
            title: freelancerData.application.profile.title,
            description: freelancerData.application.profile.description,
            photo_url: freelancerData.application.profile.photoUrl,
            hourly_charge_amount:
                freelancerData.application.hourlyChargeRate.amount,
            hourly_charge_currency:
                freelancerData.application.hourlyChargeRate.currencyCode,
            fixed_charge_amount:
                freelancerData.application.fixedChargeAmount.amount,
            fixed_charge_currency:
                freelancerData.application.fixedChargeAmount.currencyCode,
            earnings_info: this.flattenEarningsInfo(
                freelancerData.application.profile.earningsInfo
            ),
            cv: freelancerData.application.coverLetter,
            url: `https://www.upwork.com/freelancers/${freelancerData.application.profile.ciphertext}`,
            ai_reccomended: freelancerData.application.aiRecommended,
            invited: freelancerData.application.invited,
            recent_hours: freelancerData.application.profile.recentHours,
            total_hours: freelancerData.application.profile.totalHours,
            total_portfolio_items:
                freelancerData.application.profile.totalPortfolioItems,
            total_portfolio_v2_items:
                freelancerData.application.profile.totalPortfolioV2Items,
            total_feedback: freelancerData.application.profile.totalFeedback,
            recent_feedback: freelancerData.application.profile.recentFeedback,
            top_rated_status:
                freelancerData.application.profile.topRatedStatus ==
                "not_eligible"
                    ? false
                    : true,
            top_rated_plus_status:
                freelancerData.application.profile.topRatedPlusStatus ==
                "not_eligible"
                    ? false
                    : true,
            sponsored: freelancerData.application.sponsored,
            job_success_score:
                freelancerData.application.profile.jobSuccessScore,
            reccomended: freelancerData.application.recommended,
            skills: freelancerData.application.profile.skills.map(
                (skill) => skill.prettyName
            ),
            attachements: freelancerData.application.attachments.map(
                (attachment) => {
                    return {
                        name: attachment.name,
                        link: `https://www.upwork.com${attachment.link}`,
                    };
                }
            ),
            work_history: null,
        };
        return scraped_freelancer_data;
    }

    private filterLocalStorage() {
        // This regular expression matches 'modal' followed by digits, then '-expire-', followed by more digits
        const regex = /^modal\d+-expire-\d+$/;
        const matchingValues = [];
        // Access all keys in local storage
        for (let i = 0; i < localStorage.length; i++) {
            const key = localStorage.key(i);
            if (regex.test(key)) {
                // If the key matches the pattern, get the value from local storage
                const value = localStorage.getItem(key);
                matchingValues.push(JSON.parse(value));
            }
        }
        return matchingValues;
    }

    private async scrapeFreelancerCount() {
        // Select all span elements that could potentially contain the proposal count
        const spans = await this.locator.locateAll("span.air3-tab-btn-text");

        // Initialize proposalCount to null in case no matching element is found
        let proposalCount = null;

        // Iterate over the span elements to find the one that contains proposal information
        spans.forEach((span) => {
            if (span.textContent.includes("Proposals")) {
                // Extract the number from the text content of the span element
                const matches = span.textContent.match(/\d+/); // This regex matches one or more digits
                if (matches) {
                    proposalCount = parseInt(matches[0], 10); // Convert the first matched number to an integer
                }
            }
        });

        // Return the proposal count or throw an error if it wasn't found
        if (proposalCount !== null) {
            return proposalCount;
        } else {
            throw new Error("Proposal count not found.");
        }
    }

    private parseHireString(input: string): {
        total_hires: number;
        active_hire: number;
    } {
        const regex = /(\d+)\s*Hires?\s*(\d+)\s*Active/;
        const match = input.match(regex);

        if (!match) {
            throw new Error("Input string is not in the expected format");
        }

        // Extract the numbers from the match result
        const totalHires = parseInt(match[1], 10);
        const activeHire = parseInt(match[2], 10);

        // Return the result as an object
        return {
            total_hires: totalHires,
            active_hire: activeHire,
        };
    }

    private flattenEarningsInfo(earningsInfo): Earnings_Info {
        const flattened = {};
        for (const key in earningsInfo) {
            flattened[key] = earningsInfo[key].value;
        }
        return flattened as Earnings_Info;
    }

    //will try to extract the first numerical value after the first  $ in a string, parsed as a float
    private extractAmount(input: string): number | null {
        const match = input.match(/\$([0-9,]+(\.[0-9]+)?)/);
        if (match && match[1]) {
            // Remove commas and parse the number as a float
            const amount = parseFloat(match[1].replace(/,/g, ""));
            return isNaN(amount) ? null : amount;
        }
        return null;
    }
}
