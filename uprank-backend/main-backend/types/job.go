package types

import (
	"github.com/google/uuid"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent"
)

//issue is: CreateJobRequest takes in a AttachPlatformSpecificJobsRequest, which requires Job_id.

type CreateJobRequest struct {
	User_id             string                         `json:"user_id,omitempty"`
	Origin              string                         `json:"origin,omitempty"`
	PlatformJobRequests AttachPlatformSpecificJobsData `json:"platform_job_requests"`
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
	User_id                 string                         `json:"user_id,omitempty"`
	Job_id                  uuid.UUID                      `json:"job_id,omitempty"`
	PlatformSpecificJobdata AttachPlatformSpecificJobsData `json:"upwork_request,omitempty"`
}

type AttachPlatformSpecificJobsData struct {
	UpworkRequest *AttachUpworkJobData `json:"upwork_request,omitempty"`
}

func (req *AttachPlatformSpecificJobsData) Validate() map[string]interface{} {
	errors := make(map[string]interface{})

	nil_array := findNilFields(req)
	if len(nil_array) == getNumFields(req) {
		errors["PlatformJobRequests"] = "At least one platform request is required"
	}

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

type AddJobRankingRequest struct {
	Freelancer_ranking_data []FreelancerRankingData `json:"freelancer_ranking_data,omitempty"`
}

type FreelancerRankingData struct {
	Freelancer_id               string  `json:"freelancer_id,omitempty"`
	Finalized_rating_score      float32 `json:"finalized_rating_score,omitempty"`
	Raw_rating_score            float32 `json:"raw_rating_score,omitempty"`
	Uprank_reccomended          bool    `json:"uprank_reccomended,omitempty"`
	Uprank_reccomended_reasons  string  `json:"uprank_reccomended_reasons,omitempty"`
	Uprank_not_enough_data      bool    `json:"uprank_not_enough_data,omitempty"`
	Budget_adherence_percentage float32 `json:"budget_adherence_percentage,omitempty"`
	Budget_overrun_percentage   float32 `json:"budget_overrun_percentage,omitempty"`
}

func (req *FreelancerRankingData) Validate() map[string]interface{} {
	errors := make(map[string]interface{})
	if req.Finalized_rating_score < 0 {
		errors["Finalized_rating_score"] = "Finalized_rating_score must be greater than or equal to 0"
	}
	if req.Raw_rating_score < 0 {
		errors["Raw_rating_score"] = "Raw_rating_score must be greater than or equal to 0"
	}
	if req.Budget_adherence_percentage < 0 {
		errors["Budget_adherence_percentage"] = "Budget_adherence_percentage must be greater than or equal to 0"
	}

	return errors
}

func (req *AddJobRankingRequest) Validate() map[string]interface{} {
	errors := make(map[string]interface{})

	if req.Freelancer_ranking_data == nil {
		errors["Freelancer_score_map"] = "Freelancer_score_map is required"
	}
	for _, data := range req.Freelancer_ranking_data {
		data_errors := data.Validate()
		if len(data_errors) > 0 {
			errors["Freelancer_score_map"] = append(errors["Freelancer_score_map"].([]map[string]interface{}), data_errors)
		}
	}
	return errors
}

type GetAllJobsForUserRequest struct {
	User_id string `json:"user_id"`
}

type GetJobByIdRequest struct {
	Job_id  string `json:"job_id"`
	User_id string `json:"user_id"`
}
