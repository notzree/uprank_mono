package types

import "strings"

type CreateJobRequest struct {
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

func (req *CreateJobRequest) Validate() map[string]string {
	errors := make(map[string]string)

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

type ScrapeFreelancerData struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

type QueueScrapeFreelancersReqest struct {
	Job_id      string                 `json:"job_id"`
	Freelancers []ScrapeFreelancerData `json:"freelancers"`
}
