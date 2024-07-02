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

type MarkUpworkJobAsEmbeddedRequest struct {
	Job_id    string `json:"job_id,omitempty"`
	Upwork_id string `json:"upwork_id,omitempty"`
	User_id   string `json:"user_id,omitempty"`
}
type UpdateUpworkJobRequest struct {
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
