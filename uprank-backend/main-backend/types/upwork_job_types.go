package types

import (
	"strings"

	"github.com/google/uuid"
)

type AttachUpworkJobRequest struct {
	Job_Id           uuid.UUID `json:"job_id"`
	Id               string    `json:"id"`
	Title            string    `json:"title"`
	Location         string    `json:"location"`
	Description      string    `json:"description"`
	Skills           []string  `json:"skills"`
	Experience_level string    `json:"experience_level"`
	Hourly           bool      `json:"hourly"`
	Fixed            bool      `json:"fixed"`
	Hourly_rate      []float32 `json:"hourly_rate"`
	Fixed_rate       float64   `json:"fixed_rate"`
}

// Job_Id is not known from the client, is instead passed into the service from the CreateJob method.
type CreateUpworkJobRequest struct {
	Id               string    `json:"id"`
	Title            string    `json:"title"`
	Location         string    `json:"location"`
	Description      string    `json:"description"`
	Skills           []string  `json:"skills"`
	Experience_level string    `json:"experience_level"`
	Hourly           bool      `json:"hourly"`
	Fixed            bool      `json:"fixed"`
	Hourly_rate      []float32 `json:"hourly_rate"`
	Fixed_rate       float64   `json:"fixed_rate"`
}

func (req *AttachUpworkJobRequest) Validate() map[string]interface{} {
	errors := make(map[string]interface{})

	if strings.TrimSpace(req.Id) == "" {
		errors["id"] = "id cannot be nil"
	}
	if strings.TrimSpace(req.Title) == "" {
		errors["title"] = "title cannot be nil"
	}
	if strings.TrimSpace(req.Description) == "" {
		errors["description"] = "description cannot be nil"
	}

	if req.Hourly && req.Fixed {
		errors["hourly"] = "cannot have both hourly and fixed true"
		errors["fixed"] = "cannot have both hourly and fixed true"
	}
	if !req.Hourly && !req.Fixed {
		errors["hourly"] = "cannot have both hourly and fixed false"
		errors["fixed"] = "cannot have both hourly and fixed false"
	}

	return errors
}

type ScrapeUpworkFreelancerData struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}
