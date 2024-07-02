package types

import proto "github.com/notzree/uprank-backend/queue-handler/proto"

type Metadata struct {
	UserID   string `json:"user_id"`
	JobID    string `json:"job_id"`
	Type     string `json:"type"`
	Platform string `json:"platform"`
}

type UpsertVectorResponse struct {
	Job_description_vector  *proto.Vector `json:"job_description_vector"`
	Job_skill_vector        *proto.Vector `json:"job_skill_vector"`
	Upserted_freelancer_ids []string      `json:"upserted_freelancer_ids"`
}

type ComputeSpecializationRequest struct {
	Job_id                 string        `json:"job_id"`
	Work_history_count     int32         `json:"freelancer_count"`
	Job_description_vector *proto.Vector `json:"job_description_vector"`
	Job_skill_vector       *proto.Vector `json:"job_skill_vector"`
}
