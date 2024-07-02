package types

import (
	"github.com/google/uuid"
	"github.com/notzree/uprank-backend/main-backend/ent"
)

type CreateJobRequest struct {
	Origin              string                            `json:"origin,omitempty"`
	PlatformJobRequests AttachPlatformSpecificJobsRequest `json:"platform_job_requests"`
}

func (req *CreateJobRequest) Validate() map[string]interface{} {
	errors := make(map[string]interface{})

	platform_errors := req.PlatformJobRequests.Validate()
	if len(platform_errors) > 0 {
		errors["PlatformJobRequests"] = platform_errors
	}
	return errors
}

type AttachPlatformSpecificJobsRequest struct {
	UpworkRequest *AttachUpworkJobRequest `json:"upwork_request,omitempty"`
}

func (req *AttachPlatformSpecificJobsRequest) Validate() map[string]interface{} {
	errors := make(map[string]interface{})
	// nil_array := findNilFields(req)
	// if len(nil_array) == getNumFields(req) {
	// 	errors["PlatformJobRequests"] = "At least one platform request is required"
	// }
	if req.UpworkRequest != nil {
		upwork_errors := req.UpworkRequest.Validate()
		if len(upwork_errors) > 0 {
			errors["UpworkRequest"] = upwork_errors
		}
	}

	if req.UpworkRequest == nil { //When we add in more platforms, we will need to
		errors["UpworkRequest"] = "UpworkRequest is required"

	}
	return errors

}

type RankJobRequest struct {
	Job_id            uuid.UUID `json:"job_id,omitempty"`
	User_id           string    `json:"user_id,omitempty"`
	Short_lived_token string    `json:"short_lived_token,omitempty"`
	Platform          string    `json:"platform,omitempty"`
	Platform_id       string    `json:"platform_id,omitempty"`
}

type GetUpworkJobEmbeddingDataResponse struct {
	Job_id     string         `json:"job_id,omitempty"`
	Upwork_job *ent.UpworkJob `json:"upwork_job,omitempty"`
}
