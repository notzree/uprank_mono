package types

import "time"

type JobEmbeddingData struct {
	Job_id     string                  `json:"job_id,omitempty"`
	Upwork_job *UpworkJobEmbeddingData `json:"upwork_job,omitempty"`
}
type UpworkJobEmbeddingData struct {
	Upwork_id       string    `json:"id"`
	Title           string    `json:"title"`
	CreatedAt       string    `json:"created_at"`
	Location        string    `json:"location"`
	Description     string    `json:"description"`
	Skills          []string  `json:"skills"`
	ExperienceLevel string    `json:"experience_level"`
	Hourly          bool      `json:"hourly"`
	HourlyRate      []float64 `json:"hourly_rate"`
	Edges           Edges     `json:"edges"`
}

type Edges struct {
	UpworkFreelancer []Freelancer `json:"upworkfreelancer"`
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
