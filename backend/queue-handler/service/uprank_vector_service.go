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

	proto "github.com/notzree/uprank_mono/uprank-backend/queue-handler/proto"
	"github.com/notzree/uprank_mono/uprank-backend/queue-handler/types"
	sd "github.com/notzree/uprank_mono/uprank-backend/shared/service-discovery"
)

const (
	DESCRIPTION_NAMESPACE         = "description"
	SKILL_NAMESPACE               = "skill"
	WORK_HISTORY_DESCRIPTION_TYPE = "work_history_description"
	FREELANCER_SKILL_TYPE         = "freelancer_skills"
)

type UprankVecService struct {
	ServiceDiscoveryClient sd.ServiceDiscoveryClient
	MsApiKey               string
	InferenceClient        proto.InferenceClient
	HttpClient             http.Client
}

type NewUprankVecServiceInput struct {
	ServiceDiscoveryClient sd.ServiceDiscoveryClient
	MsApiKey               string
	InferenceClient        proto.InferenceClient
	HttpClient             http.Client
}

func NewUprankVecService(params NewUprankVecServiceInput) *UprankVecService {
	return &UprankVecService{
		ServiceDiscoveryClient: params.ServiceDiscoveryClient,
		MsApiKey:               params.MsApiKey,
		InferenceClient:        params.InferenceClient,
		HttpClient:             params.HttpClient,
	}
}

func (s *UprankVecService) UpsertVectors(ctx context.Context, req types.JobData, user_id string) (*types.UpsertVectorResponse, error) {
	upserted_freelancer_ids := []string{} //freelancers that have been upserted or are already upserted.

	job_description_vector, embed_err := s.InferenceClient.EmbedText(ctx, &proto.EmbedTextRequest{
		Id:       req.Upwork_job.Upwork_id,
		Text:     req.Upwork_job.Description,
		Metadata: CreateMetadata(user_id, req.Job_id, "job_description", "upwork"),
	})
	if embed_err != nil {
		return nil, embed_err
	}
	_, upsert_vector_err := s.InferenceClient.UpsertVector(ctx, &proto.UpsertVectorRequest{
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
		freelancer_description_vector, embed_description_vector_err := s.InferenceClient.EmbedText(ctx, &proto.EmbedTextRequest{
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
			work_history_description_vector, embed_work_history_vector_err := s.InferenceClient.EmbedText(ctx, &proto.EmbedTextRequest{
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
			_, upsert_vector_err := s.InferenceClient.UpsertVector(ctx, &proto.UpsertVectorRequest{
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

func (s *UprankVecService) ComputeRawSpecializationScore(ctx context.Context, req types.ComputeRawSpecializationScoreRequest) (*types.ComputeRawSpecializationScoreResponse, error) {
	upwork_job_description_vector := req.Job_description_vector
	description_scores := make(map[string]map[int]float32) //map of freelancer ids: array of the similarity scores of their previously worked jobs
	description_filter := make(map[string]string)
	description_filter["job_id"] = req.Job_data.Job_id
	description_filter["type"] = WORK_HISTORY_DESCRIPTION_TYPE
	work_history_count := int32(s.CountTotalWorkHistories(req.Job_data))
	description_response, err := s.InferenceClient.QueryVector(ctx, &proto.QueryVectorRequest{
		Namespace: DESCRIPTION_NAMESPACE,
		Vector:    upwork_job_description_vector.Vector,
		TopK:      work_history_count,
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
type DescriptionWeight func(job_data types.JobData, score_data map[string]map[int]float32) map[string]map[int]float32

func (s *UprankVecService) ApplySpecializationScoreWeights(req types.ApplySpecializationScoreWeightsRequest, ctx context.Context, weights ...DescriptionWeight) (*types.ApplySpecializationScoreWeightsResponse, error) {
	new_weights := make(map[string]map[int]float32)
	freelancers := req.Job_data.Upwork_job.Edges.UpworkFreelancer
	score_data := req.Description_scores
	remaining_job_data_map := make(map[string]types.FreelancerRankingData)
	for _, remaining_data := range req.Remaining_ranking_data {
		remaining_job_data_map[remaining_data.Freelancer_id] = remaining_data
	}

	const (
		SIMILARITY_SCORE_WEIGHT = 0.7
		BUDGET_ADHERENCE_WEIGHT = 0.22
		BUDGET_OVERRUN_WEIGHT   = 0.08
	)
	for _, freelancer := range freelancers {
		if _, exists := score_data[freelancer.ID]; exists {
			work_histories := freelancer.Edges.WorkHistories
			for _, work_history := range work_histories {
				if remaining_data, remaining_data_exists := remaining_job_data_map[freelancer.ID]; remaining_data_exists {
					if score, score_exists := score_data[freelancer.ID][work_history.ID]; score_exists {
						job_similarity_score := ExponentialScaling(score)
						var weighted_combined_budget_score float32

						// Normalize the budget adherence percentage
						normalized_budget_adherence_score := remaining_data.Budget_adherence_percentage / 100
						// Normalize the budget overrun score
						normalized_budget_overrun_score := 1 / (1 + normalized_budget_adherence_score)

						// Calculate the weighted combined budget score
						weighted_combined_budget_score = normalized_budget_overrun_score*BUDGET_OVERRUN_WEIGHT +
							normalized_budget_adherence_score*BUDGET_ADHERENCE_WEIGHT

						// Combine the similarity score with the weighted budget score
						new_weight := SIMILARITY_SCORE_WEIGHT*job_similarity_score + weighted_combined_budget_score

						// Initialize the map if not already initialized
						if new_weights[freelancer.ID] == nil {
							new_weights[freelancer.ID] = make(map[int]float32)
						}
						// Assign the new weight
						new_weights[freelancer.ID][work_history.ID] = new_weight
					}
				}
			}
		}
	}
	return &types.ApplySpecializationScoreWeightsResponse{
		Weighted_scores: new_weights,
	}, nil
}

func (s *UprankVecService) SaveRawSpecializationScoreWeights(ctx context.Context, req *types.ComputeRawSpecializationScoreResponse, data *[]types.FreelancerRankingData) error {
	summed_scores := make(map[string]float32)
	for freelancer_id, scores := range *req.Job_description_specialization_scores {
		sum := float32(0)
		for _, score := range scores {
			sum += score
		}
		average_score := sum / float32(len(scores))
		summed_scores[freelancer_id] = average_score
	}
	for i, freelancer := range *data {
		if score, exists := summed_scores[freelancer.Freelancer_id]; exists {
			// Access the original slice by index to update its value
			(*data)[i].Raw_rating_score = score
		}
	}
	return nil
}

func (s *UprankVecService) SaveWeightedSpecializationScoreWeights(ctx context.Context, req *types.ApplySpecializationScoreWeightsResponse, data *[]types.FreelancerRankingData) error {
	summed_scores := make(map[string]float32)
	for freelancer_id, scores := range req.Weighted_scores {
		sum := float32(0)
		for _, score := range scores {
			sum += score
		}
		average_score := sum / float32(len(scores))
		summed_scores[freelancer_id] = average_score
	}
	for i, freelancer := range *data {
		if score, exists := summed_scores[freelancer.Freelancer_id]; exists {
			// Access the original slice by index to update its value
			(*data)[i].Finalized_rating_score = score
		}
	}
	return nil
}

func ExponentialScaling(score float32) float32 {
	base := float32(2.0)
	return float32(math.Exp(float64(score*base)) / math.Exp(float64(base)))
}

func (s *UprankVecService) ApplyBudgetScores(ctx context.Context, req types.JobData, data *[]types.FreelancerRankingData) error {
	//Adds the Budget Adherence and Budget overrun percentage to the data
	scores := make(map[string]float32)
	overrunPercentages := make(map[string]float32)
	freelancers := req.Upwork_job.Edges.UpworkFreelancer
	for _, freelancer := range freelancers {
		work_histories := freelancer.Edges.WorkHistories
		work_histories_with_budget := float32(0)
		work_histories_within_budget := float32(0)
		totalOverrunPercentage := float32(0)
		overrunCount := float32(0)
		for _, work_history := range work_histories {
			if work_history.Budget == 0 {
				continue
			}
			work_histories_with_budget++
			if work_history.FreelancerEarnings <= work_history.Budget {
				work_histories_within_budget++
			} else {
				overrunCount++
				overrunPercentage := float32(work_history.FreelancerEarnings-work_history.Budget) / float32(work_history.Budget) * 100
				totalOverrunPercentage += overrunPercentage
			}
		}
		if work_histories_with_budget > 0 {
			budget_adherence_percentage := float32(work_histories_within_budget) / float32(work_histories_with_budget) * 100
			scores[freelancer.ID] = budget_adherence_percentage

			if overrunCount > 0 {
				averageOverrunPercentage := totalOverrunPercentage / float32(overrunCount)
				overrunPercentages[freelancer.ID] = averageOverrunPercentage
			} else {
				overrunPercentages[freelancer.ID] = 0
			}
		}
	}

	for i, freelancer := range *data {
		if score, exists := scores[freelancer.Freelancer_id]; exists {
			// Access the original slice by index to update its value
			(*data)[i].Budget_adherence_percentage = score
		}
		if overrunPercentage, exists := overrunPercentages[freelancer.Freelancer_id]; exists {
			(*data)[i].Budget_overrun_percentage = overrunPercentage
		}
	}
	return nil
}

func (s *UprankVecService) PostJobRankingData(req types.PostJobRankingDataRequest, ctx context.Context) error {
	base_url, err := s.ServiceDiscoveryClient.GetInstanceUrl(ctx, sd.GetInstanceUrlInput{
		ServiceName: "main-backend",
	})
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s/v1/private/jobs/%s/%s/%s/rank", *base_url, req.Job_id, req.Platform, req.Platform_id)
	bodyData := types.AddJobRankingRequest{
		Freelancer_ranking_data: req.Freelancer_ranking_data,
	}
	body, err := json.Marshal(bodyData)
	if err != nil {
		return err
	}

	httpreq, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	httpreq.Header.Set("X-API-KEY", s.MsApiKey)
	httpreq.Header.Set("USER-ID", req.User_id)
	httpreq.Header.Set("Content-Type", "application/json")
	resp, err := s.HttpClient.Do(httpreq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil

}

func (s *UprankVecService) FetchJobData(ctx context.Context, req types.UpworkRankingMessage) (*types.JobData, []types.FreelancerRankingData, error) {
	base_url, err := s.ServiceDiscoveryClient.GetInstanceUrl(ctx, sd.GetInstanceUrlInput{
		ServiceName: "main-backend",
	})
	if err != nil {
		return nil, nil, err
	}
	url := fmt.Sprintf("%s/v1/private/jobs/%s/%s/%s/embeddings/job_data", *base_url, req.Job_id, req.Platform, req.Platform_id)
	log.Println("Fetching data from:", url)

	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	httpreq, err := http.NewRequestWithContext(ctxWithTimeout, "GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	httpreq.Header.Set("X-API-KEY", s.MsApiKey)
	httpreq.Header.Set("USER-ID", req.User_id)
	resp, err := s.HttpClient.Do(httpreq)
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("failed to fetch %s | status code: %d ", url, resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	var job_data types.JobData
	err = json.Unmarshal((body), &job_data)
	if err != nil {
		return nil, nil, err
	}
	result := make([]types.FreelancerRankingData, 0)
	for _, freelancer := range job_data.Upwork_job.Edges.UpworkFreelancer {
		result = append(result, types.FreelancerRankingData{
			Freelancer_id:               freelancer.ID,
			Finalized_rating_score:      0.0,
			Raw_rating_score:            0.0,
			Uprank_reccomended:          false,
			Uprank_reccomended_reasons:  "",
			Uprank_not_enough_data:      false,
			Budget_adherence_percentage: 0.0,
		})
	}
	return &job_data, result, nil
}

func (s *UprankVecService) MarkFreelancersAsEmbedded(ctx context.Context, req types.MarkFreelancersAsEmbeddedRequest) error {
	log.Println("Marking freelancers as embedded")
	base_url, err := s.ServiceDiscoveryClient.GetInstanceUrl(ctx, sd.GetInstanceUrlInput{
		ServiceName: "main-backend",
	})
	if err != nil {
		return err
	}
	update_freelancer_url := fmt.Sprintf("%s/v1/private/jobs/%s/%s/%s/freelancers/update", *base_url, req.Job_id, "upwork", req.Upwork_job_id)

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
	httpreq.Header.Set("X-API-KEY", s.MsApiKey)
	httpreq.Header.Set("USER-ID", req.User_id)
	httpreq.Header.Set("Content-Type", "application/json")

	resp, err := s.HttpClient.Do(httpreq)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch %s | status code: %d ", update_freelancer_url, resp.StatusCode)
	}
	defer resp.Body.Close()
	return nil
}

func (s *UprankVecService) MarkUpworkJobAsEmbedded(ctx context.Context, req types.MarkUpworkJobAsEmbeddedRequest) error {
	log.Println("Marking job as embedded")
	base_url, err := s.ServiceDiscoveryClient.GetInstanceUrl(ctx, sd.GetInstanceUrlInput{
		ServiceName: "main-backend",
	})
	if err != nil {
		return err
	}
	update_freelancer_url := fmt.Sprintf("%s/v1/private/jobs/%s/%s/%s/update", *base_url, req.Job_id, "upwork", req.Upwork_id)
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
	httpreq.Header.Set("X-API-KEY", s.MsApiKey)
	httpreq.Header.Set("USER-ID", req.User_id)
	httpreq.Header.Set("Content-Type", "application/json")

	resp, err := s.HttpClient.Do(httpreq)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch %s | status code: %d ", update_freelancer_url, resp.StatusCode)
	}
	defer resp.Body.Close()
	return nil
}

func (s *UprankVecService) CountTotalWorkHistories(req types.JobData) int32 {
	total_workhistories := 0
	for _, freelancer := range req.Upwork_job.Edges.UpworkFreelancer {
		total_workhistories += len(freelancer.Edges.WorkHistories)
	}
	return int32(total_workhistories)
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

//Refactor condiserations:

//The service will be refactored
//It will become a Ranker struct, which will implement all the logic to rank a job ( all of the existing methods)
//it will expose a run method for now, which will call all the methods in order, and allocating goroutines
//in the ranking_handler function, we will instantiate a Ranker, and call the run method
//We will build this following CSP, but in the future try to move to actors
