package types

import "time"

type MarkFreelancersAsEmbeddedRequest struct {
	Job_id                  string   `json:"job_id"`
	Upwork_job_id           string   `json:"upwork_job_id"`
	User_id                 string   `json:"user_id"`
	Upserted_freelancer_ids []string `json:"upserted_freelancer_ids"`
}

type UpdateUpworkFreelancerRequest struct {
	Name                     *string                              `json:"name,omitempty"`
	Title                    *string                              `json:"title,omitempty"`
	Description              *string                              `json:"description,omitempty"`
	Location                 *ScrapedUpworkFreelancerLocationData `json:"location,omitempty"`
	Cv                       *string                              `json:"cv,omitempty"`
	Url                      string                               `json:"url,omitempty"`
	Ai_reccomended           *bool                                `json:"ai_reccomended,omitempty"`
	Fixed_charge_amount      *string                              `json:"fixed_charge_amount,omitempty"`
	Fixed_charge_currency    *string                              `json:"fixed_charge_currency,omitempty"`
	Hourly_charge_amount     *string                              `json:"hourly_charge_amount,omitempty"`
	Hourly_charge_currency   *string                              `json:"hourly_charge_currency,omitempty"`
	Invited                  *bool                                `json:"invited,omitempty"`
	Photo_url                *string                              `json:"photo_url,omitempty"`
	Recent_hours             *float64                             `json:"recent_hours,omitempty"`
	Total_hours              *float64                             `json:"total_hours,omitempty"`
	Total_portfolio_items    *int                                 `json:"total_portfolio_items,omitempty"`
	Total_portfolio_v2_items *int                                 `json:"total_portfolio_v2_items,omitempty"`
	Total_feedback           *float64                             `json:"total_feedback,omitempty"`
	Recent_feedback          *float64                             `json:"recent_feedback,omitempty"`
	Top_rated_status         *bool                                `json:"top_rated_status,omitempty"`
	Top_rated_plus_status    *bool                                `json:"top_rated_plus_status,omitempty"`
	Sponsored                *bool                                `json:"sponsored,omitempty"`
	Job_success_score        *float64                             `json:"job_success_score,omitempty"`
	Reccomended              *bool                                `json:"reccomended,omitempty"`
	Skills                   *[]string                            `json:"skills,omitempty"`
	Earnings_info            *ScrapedUpworkFreelancerEarningsData `json:"earnings_info,omitempty"`
	Attachements             *[]ScrapedUpworkAttachementData      `json:"attachements,omitempty"`
	Work_history             *[]ScrapedUpworkWorkHistoryData      `json:"work_history,omitempty"`
	Embedded_at              *time.Time                           `json:"embedded_at,omitempty"`
}

type ScrapedUpworkFreelancerLocationData struct {
	City     string `json:"city"`
	Country  string `json:"country"`
	Timezone string `json:"timezone"`
}

type ScrapedUpworkFreelancerEarningsData struct {
	Average_recent_earnings          float64 `json:"averageRecentEarnings"`
	Combined_average_recent_earnings float64 `json:"combinedAverageRecentEarnings"`
	Combined_recent_earnings         float64 `json:"combinedRecentEarnings"`
	Combined_total_earnings          float64 `json:"combinedTotalEarnings"`
	Combined_total_revenue           float64 `json:"combinedTotalRevenue"`
	Recent_earnings                  float64 `json:"recentEarnings"`
	Total_revenue                    float64 `json:"totalRevenue"`
}

type ScrapedUpworkAttachementData struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type ScrapedUpworkWorkHistoryData struct {
	Title               string  `json:"title"`
	Start_Date          string  `json:"start_date"`
	End_Date            string  `json:"end_date"`
	Description         string  `json:"description"`
	Budget              float64 `json:"budget"`
	Total_Earned        float64 `json:"total_earned"`
	Client_Total_Spend  float64 `json:"client_total_spend"`
	Client_Total_Hires  int     `json:"client_total_hires"`
	Client_Active_Hires int     `json:"client_active_hires"`
	Client_Feedback     string  `json:"client_feedback"`
	Client_Rating       float64 `json:"client_rating"`
	Client_Location     string  `json:"client_location"`
}

type Freelancer struct {
	ID                            string             `json:"id"`
	Name                          string             `json:"name"`
	Title                         string             `json:"title"`
	Description                   string             `json:"description"`
	City                          string             `json:"city"`
	Country                       string             `json:"country"`
	Timezone                      string             `json:"timezone"`
	CV                            string             `json:"cv"`
	AIRecommended                 bool               `json:"ai_recommended"`
	FixedChargeAmount             float64            `json:"fixed_charge_amount,omitempty"`
	FixedChargeCurrency           string             `json:"fixed_charge_currency"`
	HourlyChargeAmount            float64            `json:"hourly_charge_amount,omitempty"`
	HourlyChargeCurrency          string             `json:"hourly_charge_currency"`
	Invited                       bool               `json:"invited"`
	PhotoURL                      string             `json:"photo_url"`
	RecentHours                   float64            `json:"recent_hours"`
	TotalHours                    float64            `json:"total_hours"`
	CreatedAt                     time.Time          `json:"created_at"`
	UpdatedAt                     time.Time          `json:"updated_at"`
	TotalPortfolioItems           int                `json:"total_portfolio_items"`
	TotalPortfolioV2Items         int                `json:"total_portfolio_v2_items"`
	UpworkTotalFeedback           float64            `json:"upwork_total_feedback"`
	UpworkRecentFeedback          float64            `json:"upwork_recent_feedback"`
	UpworkTopRatedStatus          bool               `json:"upwork_top_rated_status"`
	UpworkTopRatedPlusStatus      bool               `json:"upwork_top_rated_plus_status"`
	UpworkSponsored               bool               `json:"upwork_sponsored"`
	UpworkJobSuccessScore         float64            `json:"upwork_job_success_score"`
	UpworkRecommended             bool               `json:"upwork_recommended"`
	Skills                        []string           `json:"skills"`
	AverageRecentEarnings         float64            `json:"average_recent_earnings"`
	CombinedAverageRecentEarnings float64            `json:"combined_average_recent_earnings"`
	CombinedRecentEarnings        float64            `json:"combined_recent_earnings"`
	CombinedTotalEarnings         float64            `json:"combined_total_earnings"`
	CombinedTotalRevenue          float64            `json:"combined_total_revenue"`
	RecentEarnings                float64            `json:"recent_earnings"`
	TotalRevenue                  float64            `json:"total_revenue"`
	UprankScore                   int                `json:"uprank_score"`
	EmbeddedAt                    *time.Time         `json:"embedded_at,omitempty"`
	UprankRecommended             bool               `json:"uprank_recommended"`
	UprankRecommendedReasons      string             `json:"uprank_recommended_reasons,omitempty"`
	UprankNotEnoughData           bool               `json:"uprank_not_enough_data"`
	Edges                         WorkHistoriesEdges `json:"edges"`
}

type WorkHistoriesEdges struct {
	WorkHistories []WorkHistory `json:"work_histories"`
}

type WorkHistory struct {
	ID                          int        `json:"id"`
	EmbeddedAt                  *time.Time `json:"embedded_at,omitempty"`
	CreatedAt                   time.Time  `json:"created_at"`
	UpdatedAt                   time.Time  `json:"updated_at"`
	Title                       string     `json:"title"`
	ClientFeedback              string     `json:"client_feedback,omitempty"`
	OverallRating               float64    `json:"overall_rating,omitempty"`
	FreelancerEarnings          float64    `json:"freelancer_earnings,omitempty"`
	StartDate                   *time.Time `json:"start_date,omitempty"`
	EndDate                     *time.Time `json:"end_date,omitempty"`
	Description                 string     `json:"description,omitempty"`
	Budget                      float64    `json:"budget,omitempty"`
	ClientRating                float64    `json:"client_rating,omitempty"`
	ClientReviewCount           int        `json:"client_review_count,omitempty"`
	ClientCountry               string     `json:"client_country,omitempty"`
	ClientTotalJobsPosted       int        `json:"client_total_jobs_posted,omitempty"`
	ClientTotalSpend            float64    `json:"client_total_spend,omitempty"`
	ClientTotalHires            int        `json:"client_total_hires,omitempty"`
	ClientActiveHires           int        `json:"client_active_hires,omitempty"`
	ClientTotalPaidHours        int        `json:"client_total_paid_hours,omitempty"`
	ClientAverageHourlyRatePaid float64    `json:"client_average_hourly_rate_paid,omitempty"`
	ClientCompanyCategory       string     `json:"client_company_category,omitempty"`
	ClientCompanySize           string     `json:"client_company_size,omitempty"`
	TotalProposals              int        `json:"total_proposals,omitempty"`
	NumberOfInterviews          int        `json:"number_of_interviews,omitempty"`
	Skills                      []string   `json:"skills,omitempty"`
}
