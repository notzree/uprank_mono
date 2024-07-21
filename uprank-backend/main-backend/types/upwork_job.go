package types

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type QueryUpworkJobRequest struct {
	Upwork_job_id string `json:"upwork_job_id,omitempty"`
	User_id       string `json:"user_id,omitempty"`
}

type AttachUpworkJobRequest struct {
	Job_id          uuid.UUID           `json:"job_id"`
	User_id         string              `json:"user_id"`
	Upwork_job_data AttachUpworkJobData `json:"upwork_job_data"`
}

type AttachUpworkJobData struct {
	Upwork_job_id    string     `json:"upwork_job_id"`
	Title            string     `json:"title"`
	Location         string     `json:"location"`
	Description      string     `json:"description"`
	Skills           []string   `json:"skills"`
	Experience_level string     `json:"experience_level"`
	Hourly           bool       `json:"hourly"`
	Fixed            bool       `json:"fixed"`
	Hourly_rate      []float32  `json:"hourly_rate"`
	Fixed_rate       float64    `json:"fixed_rate"`
	Embedded_at      *time.Time `json:"embedded_at"`
	Ranked_at        *time.Time `json:"ranked_at"`
}

func (req *AttachUpworkJobData) Validate() map[string]interface{} {
	errors := make(map[string]interface{})

	if strings.TrimSpace(req.Upwork_job_id) == "" {
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

type UpdateUpworkJobRequest struct {
	User_id string
	Data    UpdateUpworkJobData
}
type UpdateUpworkJobData struct {
	Upwork_id        string     `json:"upwork_id,omitempty"`
	Title            *string    `json:"title,omitempty"`
	Location         *string    `json:"location,omitempty"`
	Description      *string    `json:"description,omitempty"`
	Skills           *[]string  `json:"skills,omitempty"`
	Experience_level *string    `json:"experience_level,omitempty"`
	Hourly           *bool      `json:"hourly,omitempty"`
	Fixed            *bool      `json:"fixed,omitempty"`
	Hourly_rate      *[]float32 `json:"hourly_rate,omitempty"`
	Fixed_rate       *float64   `json:"fixed_rate,omitempty"`
	Embedded_at      *time.Time `json:"embedded_at,omitempty"`
	Ranked_at        *time.Time `json:"ranked_at,omitempty"`
}

func (req *UpdateUpworkJobData) Validate() map[string]interface{} {
	errors := make(map[string]interface{})
	// nilArray := findNilFields(req)
	// if len(nilArray) == getNumFields(req) {
	// 	errors["UpdateUpworkJobRequest"] = "all fields cannot be nil"
	// }

	return errors
}
