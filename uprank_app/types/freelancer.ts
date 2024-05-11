//All fragments need upwork.com/ in front of them
//These fields are taken from the data extracted from local storage
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
}
export interface Location{
    city: string;
    country: string;
    timezone: string;
}
export interface EarningsInfo {
    averageRecentEarnings: number;
    combinedAverageRecentEarnings: number;
    combinedRecentEarnings: number;
    combinedTotalEarnings: number;
    combinedTotalRevenue: number;
    recentEarnings: number;
    totalRevenue: number;
}

export interface Attachements {
    name: string;
    link: string;
}