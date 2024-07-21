package types

import (
	"strconv"
	"strings"
	"time"
)

type CreateUpworkFreelancerRequest struct {
	UserId      string
	UpworkJobId string
	Data        []CreateUpworkFreelancerData
}

type CreateUpworkFreelancerResponse struct {
	CreatedFreelancerCount int `json:"created_freelancer_count"`
	UpdatedFreelancerCount int `json:"updated_freelancer_count"`
	DeletedFreelancerCount int `json:"deleted_freelancer_count"`
}

type CreateUpworkFreelancerData struct {
	Name                     string                              `json:"name"`
	Title                    string                              `json:"title"`
	Description              string                              `json:"description"`
	Location                 ScrapedUpworkFreelancerLocationData `json:"location"`
	Cv                       string                              `json:"cv"`
	Url                      string                              `json:"url"`
	Ai_reccomended           bool                                `json:"ai_reccomended"`
	Fixed_charge_amount      string                              `json:"fixed_charge_amount"`
	Fixed_charge_currency    string                              `json:"fixed_charge_currency"`
	Hourly_charge_amount     string                              `json:"hourly_charge_amount"`
	Hourly_charge_currency   string                              `json:"hourly_charge_currency"`
	Invited                  bool                                `json:"invited"`
	Photo_url                string                              `json:"photo_url"`
	Recent_hours             float64                             `json:"recent_hours"`
	Total_hours              float64                             `json:"total_hours"`
	Total_portfolio_items    int                                 `json:"total_portfolio_items"`
	Total_portfolio_v2_items int                                 `json:"total_portfolio_v2_items"`
	Total_feedback           float64                             `json:"total_feedback"`
	Recent_feedback          float64                             `json:"recent_feedback"`
	Top_rated_status         bool                                `json:"top_rated_status"`
	Top_rated_plus_status    bool                                `json:"top_rated_plus_status"`
	Sponsored                bool                                `json:"sponsored"`
	Job_success_score        float64                             `json:"job_success_score"`
	Reccomended              bool                                `json:"reccomended"`
	Skills                   []string                            `json:"skills"`
	Earnings_info            ScrapedUpworkFreelancerEarningsData `json:"earnings_info"`
	Attachements             []ScrapedUpworkAttachementData      `json:"attachements"`
	Work_history             []ScrapedUpworkWorkHistoryData      `json:"work_history"`
	Embedded_at              time.Time                           `json:"embedded_at"`
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

func (req *CreateUpworkFreelancerData) Validate() map[string]interface{} {
	errors := make(map[string]interface{})

	if strings.TrimSpace(req.Name) == "" {
		errors["name"] = "name cannot be empty"
	}
	if strings.TrimSpace(req.Title) == "" {
		errors["title"] = "title cannot be empty"
	}
	if strings.TrimSpace(req.Description) == "" {
		errors["description"] = "description cannot be empty"
	}
	if strings.TrimSpace(req.Location.City) == "" {
		errors["location.city"] = "city cannot be empty"
	}
	if strings.TrimSpace(req.Location.Country) == "" {
		errors["location.country"] = "country cannot be empty"
	}
	if strings.TrimSpace(req.Location.Timezone) == "" {
		errors["location.timezone"] = "timezone cannot be empty"
	}
	if strings.TrimSpace(req.Cv) == "" {
		errors["cv"] = "cv cannot be empty"
	}
	if strings.TrimSpace(req.Url) == "" {
		errors["url"] = "url cannot be empty"
	}
	if req.Recent_hours < 0 {
		errors["recent_hours"] = "recent hours cannot be negative"
	}
	if req.Total_hours < 0 {
		errors["total_hours"] = "total hours cannot be negative"
	}
	if req.Total_portfolio_items < 0 {
		errors["total_portfolio_items"] = "total portfolio items cannot be negative"
	}
	if req.Total_portfolio_v2_items < 0 {
		errors["total_portfolio_v2_items"] = "total portfolio v2 items cannot be negative"
	}
	if req.Total_feedback < 0 || req.Total_feedback > 5 {
		errors["total_feedback"] = "total feedback must be between 0 and 5"
	}
	if req.Recent_feedback < 0 || req.Recent_feedback > 5 {
		errors["recent_feedback"] = "recent feedback must be between 0 and 5"
	}
	if req.Job_success_score < 0 || req.Job_success_score > 100 {
		errors["job_success_score"] = "job success score must be between 0 and 100"
	}

	// Ensure that at least one of fixed or hourly charge amount is provided
	if strings.TrimSpace(req.Fixed_charge_amount) == "" && strings.TrimSpace(req.Hourly_charge_amount) == "" {
		errors["charge_amount"] = "either fixed or hourly charge amount must be provided"
	}
	if strings.TrimSpace(req.Fixed_charge_amount) != "" {
		_, err := strconv.ParseFloat(req.Fixed_charge_amount, 64)
		if err != nil {
			errors["fixed_charge_amount"] = "fixed charge amount must be a valid float"
		}
	}

	if strings.TrimSpace(req.Hourly_charge_amount) != "" {
		_, err := strconv.ParseFloat(req.Hourly_charge_amount, 64)
		if err != nil {
			errors["hourly_charge_amount"] = "hourly charge amount must be a valid float"
		}
	}

	return errors
}

type UpsertFreelancerResponse struct {
	CreatedFreelancerCount int `json:"created_freelancer_count"`
	UpdatedFreelancerCount int `json:"updated_freelancer_count"`
	DeletedFreelancerCount int `json:"deleted_freelancer_count"`
}

type UpdateUpworkFreelancerRequest struct {
	UserId      string
	UpworkJobId string
	Data        []UpdateUpworkFreelancerData
}

type UpdateUpworkFreelancerData struct {
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

func (req *UpdateUpworkFreelancerData) Validate() map[string]interface{} {
	errors := make(map[string]interface{})
	if req.Name != nil && strings.TrimSpace(*req.Name) == "" {
		errors["name"] = "name cannot be empty"
	}
	if req.Title != nil && strings.TrimSpace(*req.Title) == "" {
		errors["title"] = "title cannot be empty"
	}
	if req.Description != nil && strings.TrimSpace(*req.Description) == "" {
		errors["description"] = "description cannot be empty"
	}
	if req.Location != nil {
		if strings.TrimSpace(req.Location.City) == "" {
			errors["location.city"] = "city cannot be empty"
		}
		if strings.TrimSpace(req.Location.Country) == "" {
			errors["location.country"] = "country cannot be empty"
		}
		if strings.TrimSpace(req.Location.Timezone) == "" {
			errors["location.timezone"] = "timezone cannot be empty"
		}
	}
	if req.Cv != nil && strings.TrimSpace(*req.Cv) == "" {
		errors["cv"] = "cv cannot be empty"
	}
	if strings.TrimSpace(req.Url) == "" {
		errors["url"] = "url cannot be empty"
	}
	if req.Recent_hours != nil && *req.Recent_hours < 0 {
		errors["recent_hours"] = "recent hours cannot be negative"
	}
	if req.Total_hours != nil && *req.Total_hours < 0 {
		errors["total_hours"] = "total hours cannot be negative"
	}
	if req.Total_portfolio_items != nil && *req.Total_portfolio_items < 0 {
		errors["total_portfolio_items"] = "total portfolio items cannot be negative"
	}
	if req.Total_portfolio_v2_items != nil && *req.Total_portfolio_v2_items < 0 {
		errors["total_portfolio_v2_items"] = "total portfolio v2 items cannot be negative"
	}
	if req.Total_feedback != nil && (*req.Total_feedback < 0 || *req.Total_feedback > 5) {
		errors["total_feedback"] = "total feedback must be between 0 and 5"
	}
	if req.Recent_feedback != nil && (*req.Recent_feedback < 0 || *req.Recent_feedback > 5) {
		errors["recent_feedback"] = "recent feedback must be between 0 and 5"
	}
	if req.Job_success_score != nil && (*req.Job_success_score < 0 || *req.Job_success_score > 100) {
		errors["job_success_score"] = "job success score must be between 0 and 100"
	}

	if req.Fixed_charge_amount != nil && strings.TrimSpace(*req.Fixed_charge_amount) != "" {
		_, err := strconv.ParseFloat(*req.Fixed_charge_amount, 64)
		if err != nil {
			errors["fixed_charge_amount"] = "fixed charge amount must be a valid float"
		}
	}
	if req.Hourly_charge_amount != nil && strings.TrimSpace(*req.Hourly_charge_amount) != "" {
		_, err := strconv.ParseFloat(*req.Hourly_charge_amount, 64)
		if err != nil {
			errors["hourly_charge_amount"] = "hourly charge amount must be a valid float"
		}
	}

	return errors
}
