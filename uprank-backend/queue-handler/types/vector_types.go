package types

import (
	proto "github.com/notzree/uprank-backend/queue-handler/proto"
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
	Job_id                 string        `json:"job_id"`
	Work_history_count     int32         `json:"work_history_count"`
	Freelancer_count       int32         `json:"freelancer_count"`
	Job_description_vector *proto.Vector `json:"job_description_vector"`
	Job_skill_vector       *proto.Vector `json:"job_skill_vector"`
}

type ComputeRawSpecializationScoreResponse struct {
	Job_description_specialization_scores *map[string]map[int]float32 `json:"job_description_specialization_scores"`
}

type ApplySpecializationScoreWeightsRequest struct {
	Description_scores map[string]map[int]float32 `json:"description_scores"`
	Job_data           JobEmbeddingData           `json:"job_data"`
}

type ApplySpecializationScoreWeightsResponse struct {
	Weighted_scores (map[string]map[int]float32) `json:"weighted_scores"`
}

type FinalizedJobRankingData struct {
	Job_id               string             `json:"job_id"`
	Platform             string             `json:"platform"`
	Platform_id          string             `json:"platform_id"`
	User_id              string             `json:"user_id"`
	Freelancer_score_map map[string]float32 `json:"freelancer_score_map"`
}

type AddJobRankingRequest struct {
	Freelancer_score_map map[string]float32 `json:"freelancer_score_map,omitempty"`
}
