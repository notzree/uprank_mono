package types

type JobDataAll struct {
	ID              string    `json:"id"`
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
	FixedChargeAmount             float64            `json:"fixed_charge_amount"`
	FixedChargeCurrency           string             `json:"fixed_charge_currency"`
	HourlyChargeAmount            float64            `json:"hourly_charge_amount"`
	HourlyChargeCurrency          string             `json:"hourly_charge_currency"`
	PhotoURL                      string             `json:"photo_url"`
	RecentHours                   float64            `json:"recent_hours"`
	TotalHours                    float64            `json:"total_hours"`
	TotalPortfolioItems           int                `json:"total_portfolio_items"`
	TotalPortfolioV2Items         int                `json:"total_portfolio_v2_items"`
	UpworkTotalFeedback           float64            `json:"upwork_total_feedback"`
	UpworkRecentFeedback          float64            `json:"upwork_recent_feedback"`
	UpworkTopRatedStatus          bool               `json:"upwork_top_rated_status"`
	UpworkJobSuccessScore         float64            `json:"upwork_job_success_score"`
	Skills                        []string           `json:"skills"`
	AverageRecentEarnings         float64            `json:"average_recent_earnings"`
	CombinedAverageRecentEarnings float64            `json:"combined_average_recent_earnings"`
	CombinedRecentEarnings        float64            `json:"combined_recent_earnings"`
	CombinedTotalEarnings         float64            `json:"combined_total_earnings"`
	CombinedTotalRevenue          float64            `json:"combined_total_revenue"`
	RecentEarnings                float64            `json:"recent_earnings"`
	TotalRevenue                  float64            `json:"total_revenue"`
	UprankUpdatedAt               string             `json:"uprank_updated_at"`
	Edges                         WorkHistoriesEdges `json:"edges"`
}

type WorkHistoriesEdges struct {
	WorkHistories []WorkHistory `json:"work_histories"`
}

type WorkHistory struct {
	ID                 int     `json:"id"`
	Title              string  `json:"title"`
	ClientFeedback     string  `json:"client_feedback"`
	OverallRating      float64 `json:"overall_rating"`
	FreelancerEarnings float64 `json:"freelancer_earnings"`
	StartDate          string  `json:"start_date"`
	EndDate            string  `json:"end_date"`
	Description        string  `json:"description"`
	Budget             float64 `json:"budget"`
	ClientCountry      string  `json:"client_country"`
	ClientTotalSpend   float64 `json:"client_total_spend"`
	ClientTotalHires   int     `json:"client_total_hires"`
	ClientActiveHires  int     `json:"client_active_hires"`
}
