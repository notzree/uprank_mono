package types

import (
	"strconv"
	"strings"
)

type CreateFreelancersRequest struct {
	Name                     string                        `json:"name"`
	Title                    string                        `json:"title"`
	Description              string                        `json:"description"`
	Location                 ScrapedFreelancerLocationData `json:"location"`
	Cv                       string                        `json:"cv"`
	Url                      string                        `json:"url"`
	Ai_reccomended           bool                          `json:"ai_reccomended"`
	Fixed_charge_amount      string                        `json:"fixed_charge_amount"`
	Fixed_charge_currency    string                        `json:"fixed_charge_currency"`
	Hourly_charge_amount     string                        `json:"hourly_charge_amount"`
	Hourly_charge_currency   string                        `json:"hourly_charge_currency"`
	Invited                  bool                          `json:"invited"`
	Photo_url                string                        `json:"photo_url"`
	Recent_hours             float64                       `json:"recent_hours"`
	Total_hours              float64                       `json:"total_hours"`
	Total_portfolio_items    int                           `json:"total_portfolio_items"`
	Total_portfolio_v2_items int                           `json:"total_portfolio_v2_items"`
	Total_feedback           float64                       `json:"total_feedback"`
	Recent_feedback          float64                       `json:"recent_feedback"`
	Top_rated_status         bool                          `json:"top_rated_status"`
	Top_rated_plus_status    bool                          `json:"top_rated_plus_status"`
	Sponsored                bool                          `json:"sponsored"`
	Job_success_score        float64                       `json:"job_success_score"`
	Reccomended              bool                          `json:"reccomended"`
	Skills                   []string                      `json:"skills"`
	Earnings_info            ScrapedFreelancerEarningsData `json:"earnings_info"`
	Attachements             []ScrapedAttachementData      `json:"attachements"`
}

type ScrapedFreelancerLocationData struct {
	City     string `json:"city"`
	Country  string `json:"country"`
	Timezone string `json:"timezone"`
}

type ScrapedFreelancerEarningsData struct {
	Average_recent_earnings          float64 `json:"average_recent_earnings"`
	Combined_average_recent_earnings float64 `json:"combined_average_recent_earnings"`
	Combined_recent_earnings         float64 `json:"combined_recent_earnings"`
	Combined_total_earnings          float64 `json:"combined_total_earnings"`
	Combined_total_revenue           float64 `json:"combined_total_revenue"`
	Recent_earnings                  float64 `json:"recent_eaernings"`
	Total_revenue                    float64 `json:"total_revenue"`
}
type ScrapedAttachementData struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

func (req *CreateFreelancersRequest) Validate() map[string]string {
	errors := make(map[string]string)

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