import { Decimal } from "@prisma/client/runtime/library";

//This is the backend DB Schema
type Upwork_Freelancer_Proposal = {
    url: string;
    name: string;
    title: string;
    description: string;
    city: string;
    country: string;
    timezone: string;
    cv: string;
    ai_reccomended: boolean;
    fixed_charge_amount: Decimal | null;
    fixed_charge_currency: string;
    hourly_charge_amount: Decimal;
    hourly_charge_currency: string;
    invited: boolean;
    photo_url: string;
    recent_hours: number;
    total_hours: number;
    total_portfolio_items: number;
    total_portfolio_v2_items: number;
    upwork_total_feedback: Decimal;
    upwork_recent_feedback: Decimal;
    upwork_top_rated_status: boolean;
    upwork_top_rated_plus_status: boolean;
    upwork_sponsored: boolean;
    upwork_job_success_score: Decimal;
    upwork_reccomended: boolean;
    skills: string[];
    average_recent_earnings: Decimal;
    combined_average_recent_earnings: Decimal;
    combined_recent_earnings: Decimal;
    combined_total_earnings: Decimal;
    combined_total_revenue: Decimal;
    recent_earnings: Decimal;
    total_revenue: Decimal;
    uprank_score: Decimal | null;
    uprank_updated_at: Date | null;
    uprank_reccomended: boolean | null;
    uprank_reccomended_reasons: string | null;
    uprank_not_enough_data: boolean;
    job_id: string;
};

//This is what gets scraped and sent to the backend
export interface Scraped_Freelancer_Data {
    name: string;
    title: string;
    description: string;
    location: Location;
    cv: string;
    url: string;
    ai_reccomended: boolean;
    fixed_charge_amount: string;
    fixed_charge_currency: string;
    hourly_charge_amount: string;
    hourly_charge_currency: string;
    invited: boolean;
    photo_url: string;
    recent_hours: number;
    total_hours: number;
    total_portfolio_items: number;
    total_portfolio_v2_items: number;
    total_feedback: Decimal;
    recent_feedback: Decimal;
    top_rated_status: boolean;
    top_rated_plus_status: boolean;
    sponsored: boolean;
    job_success_score: Decimal;
    reccomended: boolean;
    skills: string[];
    earnings_info: Earnings_Info;
    attachements: Attachements[];
};

export interface Location {
    city: string;
    country: string;
    timezone: string;
};

export interface Earnings_Info {
    average_recent_earnings: Decimal;
    combined_average_recent_earnings: Decimal;
    combined_recent_earnings: Decimal;
    combined_total_earnings: Decimal;
    combined_total_revenue: Decimal;
    recent_earnings: Decimal;
    total_revenue: Decimal;
};

export interface Attachements {
    name: string;
    link: string;
};

export interface Unstable_Scraped_Freelancer_Data {
    freelancers: Scraped_Freelancer_Data[];
    missing_fields: boolean;
    missing_freelancers: number;
};

