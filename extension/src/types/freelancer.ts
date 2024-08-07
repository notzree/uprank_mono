//This is the backend DB Schema
export type CreateUpworkFreelancerResponse = {
  url: string
  name: string
  title: string
  description: string
  city: string
  country: string
  timezone: string
  cv: string
  ai_reccomended: boolean
  fixed_charge_amount: number | null
  fixed_charge_currency: string
  hourly_charge_amount: number
  hourly_charge_currency: string
  invited: boolean
  photo_url: string
  recent_hours: number
  total_hours: number
  total_portfolio_items: number
  total_portfolio_v2_items: number
  upwork_total_feedback: number
  upwork_recent_feedback: number
  upwork_top_rated_status: boolean
  upwork_top_rated_plus_status: boolean
  upwork_sponsored: boolean
  upwork_job_success_score: number
  upwork_reccomended: boolean
  skills: string[]
  average_recent_earnings: number
  combined_average_recent_earnings: number
  combined_recent_earnings: number
  combined_total_earnings: number
  combined_total_revenue: number
  recent_earnings: number
  total_revenue: number
  uprank_score: number | null
  uprank_updated_at: Date | null
  uprank_reccomended: boolean | null
  uprank_reccomended_reasons: string | null
  uprank_not_enough_data: boolean
  job_id: string
}

//This is what gets scraped and sent to the backend
export interface Scraped_Freelancer_Data {
  name: string
  title: string
  description: string
  location: Location
  cv: string
  url: string
  ai_reccomended: boolean
  fixed_charge_amount: string
  fixed_charge_currency: string
  hourly_charge_amount: string
  hourly_charge_currency: string
  invited: boolean
  photo_url: string
  recent_hours: number
  total_hours: number
  total_portfolio_items: number
  total_portfolio_v2_items: number
  total_feedback: number
  recent_feedback: number
  top_rated_status: boolean
  top_rated_plus_status: boolean
  sponsored: boolean
  job_success_score: number
  reccomended: boolean
  skills: string[]
  earnings_info: Earnings_Info
  attachements: Attachements[]
  work_history: FreelancerJobHistory[] | null
}

export interface Location {
  city: string
  country: string
  timezone: string
}

export interface Earnings_Info {
  average_recent_earnings: number
  combined_average_recent_earnings: number
  combined_recent_earnings: number
  combined_total_earnings: number
  combined_total_revenue: number
  recent_earnings: number
  total_revenue: number
}

export interface Attachements {
  name: string
  link: string
}

export type CreateFreelancerProxyRequest = {
  update: boolean
  authentication_token: string
  freelancers: Scraped_Freelancer_Data[]
  job_id: string
}

export type CreateUpworkFreelancerClientRequest = {
    freelancers: Scraped_Freelancer_Data[]
}

export type CreateUpworkFreelancerClientResponse = {
    result: CreateUpworkFreelancerResponse | null
    error_msg: string | null;
}

export type UpdateUpworkFreelancerClientResponse = {
    result: UpdateUpworkFreelancerResponse | null
    error_msg: string | null;
}
type UpdateUpworkFreelancerResponse = {
  created_freelancer_count: number
  updated_freelancer_count: number
  deleted_freelancer_count: number
}

export interface ScrapeFreelancerResponse {
  freelancers: Scraped_Freelancer_Data[]
  missing_fields: boolean
  missing_freelancers: number
}


export interface CreateFreelancerResponse {
  ok: boolean
  count: number
}


export type FreelancerJobHistory = {
  title: string;
  start_date: string | null;
  end_date: string | null;
  description: string | null;
  budget: number | null;
  total_earned: number | null;
  client_total_spend: number | null;
  client_total_hires: number | null;
  client_active_hires: number | null;
  client_feedback: string | null;
  client_rating: number | null;
  client_location: string | null;
}

export type ProcessFreelancerJobHistoryResult = {
  name: string;
  jobs: FreelancerJobHistory[];
}