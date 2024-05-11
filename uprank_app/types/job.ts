export interface ScrapedJobData {
    id: string;
    title: string;
    description: string;
    location: string;
    hourly: boolean;
    fixed: boolean;
    hourly_rate: number[];
    fixed_rate: number[];
    num_months: number[];
    experience_level: string;
    skills: string[];
}
