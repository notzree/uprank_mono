//This is the backend DB Schema
type UpworkFreelancerProposal = {
    url: string;
    name: string;
    title: string;
    description: string;
    city: string;
    country: string;
    timezone: string;
    cv: string;
    aiReccomended: boolean;
    fixedChargeAmount: number;
    fixedChargeCurrency: string;
    hourlyChargeAmount: number;
    hourlyChargeCurrency: string;
    invited: boolean;
    photoUrl: string;
    recentHours: number;
    totalHours: number;
    totalPortfolioItems: number;
    totalPortfolioV2Items: number;
    upwork_totalFeedback: number; 
    upwork_recentFeedback: number; 
    upwork_topRatedStatus: boolean;
    upwork_topRatedPlusStatus: boolean;
    upwork_sponsored: boolean;
    upwork_jobSuccessScore: number; 
    upwork_reccomended: boolean;
    skills: string[];
    averageRecentEarnings: number;
    combinedAverageRecentEarnings: number;
    combinedRecentEarnings: number;
    combinedTotalEarnings: number;
    combinedTotalRevenue: number;
    recentEarnings: number;
    totalRevenue: number;
    uprank_score: number; 
    uprank_updated_at: Date | null;
    uprank_reccomended: boolean; 
    uprank_reccomended_reasons: string | null;
    uprank_not_enough_data: boolean; 
    jobId: string; 
};
//This is what gets scraped and sent to the backend
export interface ScrapedFreelancerData {
    name: string;
    title: string;
    description: string;
    location: Location;
    cv: string;
    url: string;
    aiReccomended: boolean;
    fixedChargeAmount: string;
    fixedChargeCurrency: string;
    hourlyChargeAmount: string;
    hourlyChargeCurrency: string;
    invited: boolean;
    photoUrl: string;
    recentHours: number;
    totalHours: number;
    totalPortfolioItems: number;
    totalPortfolioV2Items: number;
    totalFeedback: number;
    recentFeedback: number;
    topRatedStatus: boolean;
    topRatedPlusStatus: boolean;
    sponsored: boolean;
    jobSuccessScore: number;
    reccomended: boolean;
    skills: string[];
    earningsInfo: EarningsInfo;
    attachements: Attachements[];
};
export interface Location{
    city: string;
    country: string;
    timezone: string;
};
export interface EarningsInfo {
    averageRecentEarnings: number;
    combinedAverageRecentEarnings: number;
    combinedRecentEarnings: number;
    combinedTotalEarnings: number;
    combinedTotalRevenue: number;
    recentEarnings: number;
    totalRevenue: number;
};

export interface Attachements {
    name: string;
    link: string;
};
export interface SendFreelancerBody {
    authentication_token: string;
    freelancers: ScrapedFreelancerData[];
    jobId: string;
};

export interface UnstableScrapedFreelancerData {
    freelancers: ScrapedFreelancerData[];
    missingFields: boolean;
    missingFreelancers: number;
};

export interface SendFreelancerResponse {
    ok: boolean;
    count: number;
};

