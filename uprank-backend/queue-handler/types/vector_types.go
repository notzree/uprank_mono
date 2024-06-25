package types

type Metadata struct {
	UserID   string `json:"user_id"`
	JobID    string `json:"job_id"`
	Type     string `json:"type"`
	Platform string `json:"platform"`
}
