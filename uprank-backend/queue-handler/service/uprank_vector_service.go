package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
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
	_, upsert_vector_err := s.infer.UpsertVector(ctx, &proto.UpsertVectorRequest{
		Namespace: DESCRIPTION_NAMESPACE,
		Vectors:   []*proto.Vector{job_description_vector.Vector},
	})
	if upsert_vector_err != nil {
		return nil, upsert_vector_err
	}

	for _, freelancer := range req.Upwork_job.Edges.UpworkFreelancer {
		if !freelancer.EmbeddedAt.IsZero() && freelancer.EmbeddedAt.Before(freelancer.UpdatedAt) {
			continue
		}
		description_vectors := []*proto.Vector{}

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

		for _, work_history := range freelancer.Edges.WorkHistories {
			if !freelancer.EmbeddedAt.IsZero() && work_history.EmbeddedAt.After(work_history.UpdatedAt) {
				continue
			}
			if work_history.Description == "" {
				continue
			}
			work_history_description_vector, embed_work_history_vector_err := s.infer.EmbedText(ctx, &proto.EmbedTextRequest{
				Id:       strconv.Itoa(work_history.ID),
				Text:     work_history.Description,
				Metadata: CreateMetadata(user_id, req.Job_id, "work_history_description", "upwork", WithFreelancerId(freelancer.ID), WithWorkHistoryId(work_history.ID)),
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
	}, nil
}

func (s *UprankVecService) ComputeRawSpecializationScore(req types.ComputeRawSpecializationScoreRequest, ctx context.Context) (*types.ComputeRawSpecializationScoreResponse, error) {
	upwork_job_description_vector := req.Job_description_vector
	description_scores := make(map[string]map[int]float32) //map of freelancer ids: array of the similarity scores of their previously worked jobs
	description_filter := make(map[string]string)
	description_filter["job_id"] = req.Job_id
	description_filter["type"] = WORK_HISTORY_DESCRIPTION_TYPE

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
		work_history_id := vector.Metadata["work_history_id"]
		work_history_id_int, err := strconv.Atoi(work_history_id)
		if err != nil {
			return nil, err
		}
		if _, exists := description_scores[vector_id]; !exists {
			description_scores[vector_id] = map[int]float32{
				work_history_id_int: vector.Score,
			}
		} else {
			description_scores[vector_id][work_history_id_int] = vector.Score
		}
	}

	return &types.ComputeRawSpecializationScoreResponse{
		Job_description_specialization_scores: &description_scores,
	}, nil
}

// Weight is a function that takes in some job data, as well as the existing scores, and returns the scores with some weighting applied.
// The job_data is used in the case that the weight is dependent on the job data, such as the job duration. budget, etc.
type DescriptionWeight func(job_data types.JobEmbeddingData, score_data map[string]map[int]float32) map[string]map[int]float32

func (s *UprankVecService) ApplySpecializationScoreWeights(req types.ApplySpecializationScoreWeightsRequest, ctx context.Context, weights ...DescriptionWeight) (*types.ApplySpecializationScoreWeightsResponse, error) {
	new_weights := make(map[string]map[int]float32)
	freelancers := req.Job_data.Upwork_job.Edges.UpworkFreelancer
	score_data := req.Description_scores
	for _, freelancer := range freelancers {
		if _, exists := score_data[freelancer.ID]; exists {
			work_histories := freelancer.Edges.WorkHistories
			for _, work_history := range work_histories {
				if _, exists := score_data[freelancer.ID][work_history.ID]; exists {
					var budget_adherence_score float32
					job_similarity_score := ExponentialScaling(score_data[freelancer.ID][work_history.ID])
					if work_history.Budget != 0 {
						if work_history.FreelancerEarnings > work_history.Budget {
							budget_adherence_score = 0.0
						} else {
							budget_adherence_score = 1.0
						}
					} else {
						budget_adherence_score = 0.5
					}
					new_weight := 0.95*job_similarity_score + 0.05*budget_adherence_score
					if new_weights[freelancer.ID] == nil {
						new_weights[freelancer.ID] = make(map[int]float32)
						new_weights[freelancer.ID][work_history.ID] = new_weight
					}
					new_weights[freelancer.ID][work_history.ID] = new_weight
				}
			}
		}

	}
	return &types.ApplySpecializationScoreWeightsResponse{
		Weighted_scores: new_weights,
	}, nil
}

func ExponentialScaling(score float32) float32 {
	base := float32(2.0)
	return float32(math.Exp(float64(score*base)) / math.Exp(float64(base)))
}

func (s *UprankVecService) PostJobRankingData(req types.FinalizedJobRankingData, ctx context.Context) error {
	url := fmt.Sprintf("%s/v1/private/jobs/%s/%s/%s/rank", s.backend_url, req.Job_id, req.Platform, req.Platform_id)
	bodyData := types.AddJobRankingRequest{
		Freelancer_score_map: req.Freelancer_score_map,
	}
	body, err := json.Marshal(bodyData)
	if err != nil {
		return err
	}

	httpreq, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
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
