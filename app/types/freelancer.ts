export type PortfolioItem = {
  title: string;
  description: string;
  url: string;
};

export type UpworkFreelancer = {
  id: string;
  name: string;
  title: string;
  description: string;
  city: string;
  country: string;
  timezone: string;
  cv: string;
  fixed_charge_amount: number;
  fixed_charge_currency: string;
  hourly_charge_amount: number;
  hourly_charge_currency: string;
  photo_url: string;
  recent_hours: number;
  total_hours: number;
  created_at: string;
  updated_at: string;
  embedded_at: string;
  total_portfolio_items: number;
  upwork_total_feedback: number;
  upwork_recent_feedback: number;
  upwork_top_rated_status: boolean;
  upwork_job_success_score: number;
  skills: string[];
  average_recent_earnings: number;
  combined_average_recent_earnings: number;
  combined_recent_earnings: number;
  combined_total_earnings: number;
  combined_total_revenue: number;
  recent_earnings: number;
  total_revenue: number;
  portfolio_items?: PortfolioItem[];
  edges: UpworkFreelancerEdges;
};

export type UpworkFreelancerEdges = {
  freelancer_inference_data: FreelancerInferenceData
}

export type FreelancerInferenceData = {
  ai_estimated_duration: any;
  finalized_rating_score: number;
  raw_rating_score: number;
  uprank_reccomended: boolean;
  uprank_reccomended_reasons: string;
  uprank_not_enough_data: boolean;
  budget_adherence_percentage: number;
  budget_overrun_percentage: number;
  id: number;
  edges: any;
}