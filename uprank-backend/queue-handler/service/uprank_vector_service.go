package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	proto "github.com/notzree/uprank-backend/queue-handler/proto"
	"github.com/notzree/uprank-backend/queue-handler/types"
)

const (
	DESCRIPTION_NAMESPACE         = "description"
	SKILL_NAMESPACE               = "skill"
	WORK_HISTORY_DESCRIPTION_TYPE = "work_history_description"
	FREELANCER_SKILL_TYPE         = "freelancer_skills"
)

type UprankVecService struct {
	backend_url string
	ms_api_key  string
	infer       proto.InferenceClient
	httpClient  http.Client
}

func NewUprankVecService(backend_url string, ms_api_key string, infer proto.InferenceClient, httpClient http.Client) *UprankVecService {
	return &UprankVecService{
		backend_url: backend_url,
		ms_api_key:  ms_api_key,
		infer:       infer,
		httpClient:  httpClient,
	}
}


func (s *UprankVecService) UpsertVectors(req types.JobEmbeddingData, user_id string) (*types.UpsertVectorResponse, error) {
	upserted_freelancer_ids := []string{} //freelancers that have been upserted or are already upserted.
	ctx := context.Background()
	job_description_vector, embed_err := s.infer.EmbedText(ctx, &proto.EmbedTextRequest{
		Id:       req.Upwork_job.Upwork_id,
		Text:     req.Upwork_job.Description,
		Metadata: CreateMetadata(user_id, req.Job_id, "job_description", "upwork"),
	})
	if embed_err != nil {
		return nil, embed_err
	}
	job_skills_as_string := strings.Join(req.Upwork_job.Skills, " ")
	job_skill_vector, embed_err := s.infer.EmbedText(ctx, &proto.EmbedTextRequest{
		Text:     job_skills_as_string,
		Id:       req.Upwork_job.Upwork_id,
		Metadata: CreateMetadata(user_id, req.Job_id, "job_skills", "upwork"),
	})
	if embed_err != nil {
		return nil, embed_err
	}
	_, upsert_vector_err := s.infer.UpsertVector(ctx, &proto.UpsertVectorRequest{
		Namespace: DESCRIPTION_NAMESPACE,
		Vectors:   []*proto.Vector{job_description_vector.Vector},
	})
	if upsert_vector_err != nil {
		return nil, upsert_vector_err
	}
	_, upsert_vector_err = s.infer.UpsertVector(ctx, &proto.UpsertVectorRequest{
		Namespace: SKILL_NAMESPACE,
		Vectors:   []*proto.Vector{job_skill_vector.Vector},
	})
	if upsert_vector_err != nil {
		return nil, upsert_vector_err
	}

	for _, freelancer := range req.Upwork_job.Edges.UpworkFreelancer {
		if freelancer.EmbeddedAt != nil && freelancer.EmbeddedAt.Before(freelancer.UpdatedAt) {
			continue
		}
		description_vectors := []*proto.Vector{}
		skill_vectors := []*proto.Vector{}

		//Embed and add vectors to array to embed in 1 trip
		freelancer_description_vector, embed_description_vector_err := s.infer.EmbedText(ctx, &proto.EmbedTextRequest{
			Id:       freelancer.ID,
			Text:     freelancer.Description,
			Metadata: CreateMetadata(user_id, req.Job_id, "freelancer_description", "upwork", WithFreelancerId(freelancer.ID)),
		})
		if embed_description_vector_err != nil {
			return nil, embed_description_vector_err
		}

		description_vectors = append(description_vectors, freelancer_description_vector.Vector)

		freelancer_skills_as_string := strings.Join(freelancer.Skills, " ")
		freelancer_skill_vector, embed_skill_vector_err := s.infer.EmbedText(ctx, &proto.EmbedTextRequest{
			Id:       freelancer.ID,
			Text:     freelancer_skills_as_string,
			Metadata: CreateMetadata(user_id, req.Job_id, "freelancer_skills", "upwork", WithFreelancerId(freelancer.ID)),
		})
		if embed_skill_vector_err != nil {
			return nil, embed_skill_vector_err
		}
		skill_vectors = append(skill_vectors, freelancer_skill_vector.Vector)

		for _, work_history := range freelancer.Edges.WorkHistories {
			if work_history.EmbeddedAt != nil && work_history.EmbeddedAt.After(work_history.UpdatedAt) {
				continue
			}
			if work_history.Description == "" {
				continue
			}
			work_history_description_vector, embed_work_history_vector_err := s.infer.EmbedText(ctx, &proto.EmbedTextRequest{
				Id:       strconv.Itoa(work_history.ID),
				Text:     work_history.Description,
				Metadata: CreateMetadata(user_id, req.Job_id, "work_history_description", "upwork", WithFreelancerId(freelancer.ID), WithWorkHistoryId(work_history.ID),
			})
			if embed_work_history_vector_err != nil {
				return nil, embed_work_history_vector_err
			}

			description_vectors = append(description_vectors, work_history_description_vector.Vector)
		}
		if len(description_vectors) != 0 {
			_, upsert_vector_err := s.infer.UpsertVector(ctx, &proto.UpsertVectorRequest{
				Namespace: DESCRIPTION_NAMESPACE,
				Vectors:   description_vectors,
			})
			if upsert_vector_err != nil {
				return nil, upsert_vector_err
			}
		}
		if len(skill_vectors) != 0 {
			_, upsert_vector_err = s.infer.UpsertVector(ctx, &proto.UpsertVectorRequest{
				Namespace: SKILL_NAMESPACE,
				Vectors:   skill_vectors,
			})
			if upsert_vector_err != nil {
				return nil, upsert_vector_err
			}
		}
		upserted_freelancer_ids = append(upserted_freelancer_ids, freelancer.ID)
	}

	//todo: Make these run concurrently!?
	marking_job_err := s.MarkUpworkJobAsEmbedded(ctx, types.MarkUpworkJobAsEmbeddedRequest{
		Job_id:    req.Job_id,
		Upwork_id: req.Upwork_job.Upwork_id,
		User_id:   user_id,
	})
	if marking_job_err != nil {
		return nil, marking_job_err
	}
	marking_freelancer_err := s.MarkFreelancersAsEmbedded(ctx, types.MarkFreelancersAsEmbeddedRequest{
		Job_id:                  req.Job_id,
		Upwork_job_id:           req.Upwork_job.Upwork_id,
		User_id:                 user_id,
		Upserted_freelancer_ids: upserted_freelancer_ids,
	})
	if marking_freelancer_err != nil {
		return nil, marking_freelancer_err
	}

	return &types.UpsertVectorResponse{
		Job_description_vector: job_description_vector.Vector,
		Job_skill_vector:       job_skill_vector.Vector,
	}, nil
}

func (s *UprankVecService) ComputeRawSpecializationScore(req types.ComputeRawSpecializationScoreRequest, ctx context.Context) (*types.ComputeRawSpecializationScoreResponse, error) {
	upwork_job_description_vector := req.Job_description_vector
	description_scores := make(map[string]map[string]float32) //map of freelancer ids: array of the similarity scores of their previously worked jobs
	description_filter := make(map[string]string)
	description_filter["job_id"] = req.Job_id
	description_filter["type"] = WORK_HISTORY_DESCRIPTION_TYPE

	upwork_job_skill_vector := req.Job_skill_vector
	skill_scores := make(map[string]float32) // map of freelancer ids: freelancer skill similarity score to job
	skill_filter := make(map[string]string)
	skill_filter["job_id"] = req.Job_id
	skill_filter["type"] = FREELANCER_SKILL_TYPE

	description_response, err := s.infer.QueryVector(ctx, &proto.QueryVectorRequest{
		Namespace: DESCRIPTION_NAMESPACE,
		Vector:    upwork_job_description_vector.Vector,
		TopK:      req.Work_history_count,
		Filter:    description_filter,
	})
	if err != nil {
		return nil, err
	}
	for _, vector := range description_response.Matches {
		vector_id := vector.Metadata["freelancer_id"]
		job_id :=vector.Metadata['']
		if _, exists := description_scores[vector_id]; !exists {
			description_scores[vector_id] = map[string]float32{
				"": vector.Score,
			}
		} else {
			description_scores[vector_id] = append(description_scores[vector_id], vector.Score)
		}
	}

	skill_response, err := s.infer.QueryVector(ctx, &proto.QueryVectorRequest{
		Namespace: SKILL_NAMESPACE,
		Vector:    upwork_job_skill_vector.Vector,
		TopK:      req.Freelancer_count,
		Filter:    description_filter,
	})
	if err != nil {
		return nil, err
	}

	for _, vector := range skill_response.Matches {
		vector_id := vector.Metadata["freelancer_id"]
		if _, exists := skill_scores[vector_id]; !exists {
			skill_scores[vector_id] = vector.Score
		} else {
			return nil, fmt.Errorf("Duplicate freelancer id in skill scores")
		}
	}

	return &types.ComputeRawSpecializationScoreResponse{
		Job_description_specialization_scores: &description_scores,
		Job_skill_specialization_scores:       &skill_scores,
	}, nil
}

func (s *UprankVecService) ApplySpecializationScoreWeights(req types.ComputeRawSpecializationScoreResponse, ctx context.Context) (*types.ApplySpecializationScoreWeightsResponse, error) {
	filePath := "score_output.json"
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(req)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *UprankVecService) FetchJobData(req types.UpworkRankingMessage) (*types.JobEmbeddingData, error) {
	fetch_url := fmt.Sprintf("%s/v1/private/jobs/%s/%s/%s/embeddings/job_data", s.backend_url, req.Job_id, req.Platform, req.Platform_id)
	log.Println("Fetching data from:", fetch_url)
	httpreq, err := http.NewRequest("GET", fetch_url, nil)
	if err != nil {
		return nil, err
	}
	httpreq.Header.Set("X-API-KEY", s.ms_api_key)
	httpreq.Header.Set("User_id", req.User_id)
	resp, err := s.httpClient.Do(httpreq)
	if err != nil {
		return nil, err

	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var job_data types.JobEmbeddingData
	err = json.Unmarshal((body), &job_data)
	if err != nil {
		return nil, err
	}
	return &job_data, nil
}

func (s *UprankVecService) MarkFreelancersAsEmbedded(ctx context.Context, req types.MarkFreelancersAsEmbeddedRequest) error {
	log.Println("Marking freelancers as embedded")
	//todo: change the date to be whenever its embedded, not when its marked as embedded
	update_freelancer_url := fmt.Sprintf("%s/v1/private/jobs/%s/%s/%s/freelancers/update", s.backend_url, req.Job_id, "upwork", req.Upwork_job_id)

	bodyData := []types.UpdateUpworkFreelancerRequest{}
	for _, freelancer_id := range req.Upserted_freelancer_ids {
		current_time := time.Now()
		bodyData = append(bodyData, types.UpdateUpworkFreelancerRequest{
			Url:         freelancer_id,
			Embedded_at: &current_time,
		})
	}
	body, err := json.Marshal(bodyData)
	if err != nil {
		return err
	}
	httpreq, err := http.NewRequest("POST", update_freelancer_url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	httpreq.Header.Set("X-API-KEY", s.ms_api_key)
	httpreq.Header.Set("User_id", req.User_id)
	httpreq.Header.Set("Content-Type", "application/json")

	resp, err := s.httpClient.Do(httpreq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (s *UprankVecService) MarkUpworkJobAsEmbedded(ctx context.Context, req types.MarkUpworkJobAsEmbeddedRequest) error {
	log.Println("Marking job as embedded")
	update_freelancer_url := fmt.Sprintf("%s/v1/private/jobs/%s/%s/%s/update", s.backend_url, req.Job_id, "upwork", req.Upwork_id)
	current_time := time.Now()
	bodyData := types.UpdateUpworkJobRequest{
		Upwork_id:   req.Upwork_id,
		Embedded_at: &current_time,
	}
	body, err := json.Marshal(bodyData)
	if err != nil {
		return err
	}
	httpreq, err := http.NewRequest("POST", update_freelancer_url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	httpreq.Header.Set("X-API-KEY", s.ms_api_key)
	httpreq.Header.Set("User_id", req.User_id)
	httpreq.Header.Set("Content-Type", "application/json")

	resp, err := s.httpClient.Do(httpreq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (s *UprankVecService) CountTotalWorkHistories(req types.JobEmbeddingData) int {
	total_workhistories := 0
	for _, freelancer := range req.Upwork_job.Edges.UpworkFreelancer {
		total_workhistories += len(freelancer.Edges.WorkHistories)
	}
	return total_workhistories
}


type MetadataOption func(map[string]string)

func WithFreelancerId(freelancer_id string) MetadataOption {
	return func(metadata map[string]string) {
		metadata["freelancer_id"] = freelancer_id
	}
}

func WithWorkHistoryId(work_history_id int) MetadataOption {
	return func(metadata map[string]string) {
		metadata["work_history_id"] = strconv.Itoa(work_history_id)
	}
}

func CreateMetadata(user_id string, job_id string, vector_type string, platform string, options ...MetadataOption) map[string]string {

	metadata := make(map[string]string)
	metadata["user_id"] = user_id
	metadata["job_id"] = job_id
	metadata["type"] = vector_type
	metadata["platform"] = platform

	for _, option := range options {
		option(metadata)
	}
	return metadata
}

// func CreateMetadata(user_id string, job_id string, vector_type string, platform string, freelancer_id *string, work_history_id *int) map[string]string {
// 	metadata := make(map[string]string)
// 	metadata["user_id"] = user_id
// 	metadata["job_id"] = job_id
// 	metadata["type"] = vector_type
// 	metadata["platform"] = platform
// 	if work_history_id != nil {
// 		metadata["work_history_id"] = strconv.Itoa(*work_history_id)

// 	}
// 	if freelancer_id != nil {
// 		metadata["freelancer_id"] = *freelancer_id
// 	}
// 	return metadata
// }
