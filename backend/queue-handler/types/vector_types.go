package types

import (
	proto "github.com/notzree/uprank_mono/uprank-backend/queue-handler/proto"
)

type Metadata struct {
	UserID   string `json:"user_id"`
	JobID    string `json:"job_id"`
	Type     string `json:"type"`
	Platform string `json:"platform"`
}

type UpsertVectorResponse struct {
	Job_description_vector *proto.Vector `json:"job_description_vector"`
}

type ComputeRawSpecializationScoreRequest struct {
	Job_data               JobData       `json:"job_data"`
	Job_description_vector *proto.Vector `json:"job_description_vector"`
	Job_skill_vector       *proto.Vector `json:"job_skill_vector"`
}

type ComputeRawSpecializationScoreResponse struct {
	Job_description_specialization_scores *map[string]map[int]float32 `json:"job_description_specialization_scores"`
}

type ApplySpecializationScoreWeightsRequest struct {
	Description_scores     map[string]map[int]float32 `json:"description_scores"`
	Remaining_ranking_data []FreelancerRankingData    `json:"remaining_ranking_data"`
	Job_data               JobData                    `json:"job_data"`
}

type ApplySpecializationScoreWeightsResponse struct {
	Weighted_scores (map[string]map[int]float32) `json:"weighted_scores"`
}

type PostJobRankingDataRequest struct {
	Job_id                  string                  `json:"job_id"`
	Platform                string                  `json:"platform"`
	Platform_id             string                  `json:"platform_id"`
	User_id                 string                  `json:"user_id"`
	Freelancer_ranking_data []FreelancerRankingData `json:"freelancer_ranking_data"`
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
