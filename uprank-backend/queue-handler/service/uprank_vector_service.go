package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	proto "github.com/notzree/uprank-backend/queue-handler/proto"
	"github.com/notzree/uprank-backend/queue-handler/types"
)

const (
	DESCRIPTION_NAMESPACE         = "description"
	SKILL_NAMESPACE               = "skill"
	WORK_HISTORY_DESCRIPTION_TYPE = "work_history_description"
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
	upserted_freelancer_ids := []string{}
	ctx := context.Background()
	job_description_vector, embed_err := s.infer.EmbedText(ctx, &proto.EmbedTextRequest{
		Id:       req.Upwork_job.Upwork_id,
		Text:     req.Upwork_job.Description,
		Metadata: CreateMetadata(user_id, req.Job_id, "job_description", "upwork", nil),
	})
	if embed_err != nil {
		return nil, embed_err
	}
	job_skills_as_string := strings.Join(req.Upwork_job.Skills, " ")
	job_skill_vector, embed_err := s.infer.EmbedText(ctx, &proto.EmbedTextRequest{
		Text:     job_skills_as_string,
		Id:       req.Upwork_job.Upwork_id,
		Metadata: CreateMetadata(user_id, req.Job_id, "job_skills", "upwork", nil),
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
		description_vectors := []*proto.Vector{}
		skill_vectors := []*proto.Vector{}

		//Embed and add vectors to array to embed in 1 trip
		freelancer_description_vector, embed_description_vector_err := s.infer.EmbedText(ctx, &proto.EmbedTextRequest{
			Id:       freelancer.ID,
			Text:     freelancer.Description,
			Metadata: CreateMetadata(user_id, req.Job_id, "freelancer_description", "upwork", &freelancer.ID),
		})
		if embed_description_vector_err != nil {
			return nil, embed_description_vector_err
		}

		description_vectors = append(description_vectors, freelancer_description_vector.Vector)

		freelancer_skills_as_string := strings.Join(freelancer.Skills, " ")
		freelancer_skill_vector, embed_skill_vector_err := s.infer.EmbedText(ctx, &proto.EmbedTextRequest{
			Id:       freelancer.ID,
			Text:     freelancer_skills_as_string,
			Metadata: CreateMetadata(user_id, req.Job_id, "freelancer_skills", "upwork", &freelancer.ID),
		})
		if embed_skill_vector_err != nil {
			return nil, embed_skill_vector_err
		}
		skill_vectors = append(skill_vectors, freelancer_skill_vector.Vector)

		for _, work_history := range freelancer.Edges.WorkHistories {
			if work_history.Description == "" {
				continue
			}
			work_history_description_vector, embed_work_history_vector_err := s.infer.EmbedText(ctx, &proto.EmbedTextRequest{
				Id:       strconv.Itoa(work_history.ID),
				Text:     work_history.Description,
				Metadata: CreateMetadata(user_id, req.Job_id, "work_history_description", "upwork", &freelancer.ID),
			})
			if embed_work_history_vector_err != nil {
				return nil, embed_work_history_vector_err
			}

			description_vectors = append(description_vectors, work_history_description_vector.Vector)
		}

		_, upsert_vector_err := s.infer.UpsertVector(ctx, &proto.UpsertVectorRequest{
			Namespace: DESCRIPTION_NAMESPACE,
			Vectors:   description_vectors,
		})
		if upsert_vector_err != nil {
			return nil, upsert_vector_err
		}
		_, upsert_vector_err = s.infer.UpsertVector(ctx, &proto.UpsertVectorRequest{
			Namespace: SKILL_NAMESPACE,
			Vectors:   skill_vectors,
		})
		if upsert_vector_err != nil {
			return nil, upsert_vector_err
		}
		upserted_freelancer_ids = append(upserted_freelancer_ids, freelancer.ID)
	}

	return &types.UpsertVectorResponse{
		Job_description_vector:  job_description_vector.Vector,
		Job_skill_vector:        job_skill_vector.Vector,
		Upserted_freelancer_ids: upserted_freelancer_ids,
	}, nil
}

func (s *UprankVecService) ComputeSpecialization(req types.ComputeSpecializationRequest, ctx context.Context) (map[string][]float32, error) {
	scores := make(map[string][]float32) //map of freelancer ids : scores + other data
	upwork_job_description_vector := req.Job_description_vector
	// upwork_job_skill_vector := req.Job_skill_vector
	//given the above vectors, try to find the best freelancers
	// number_of_freelancers := len(req.Upserted_freelancer_ids)

	filter := make(map[string]string)
	filter["job_id"] = req.Job_id
	filter["type"] = WORK_HISTORY_DESCRIPTION_TYPE
	response, err := s.infer.QueryVector(ctx, &proto.QueryVectorRequest{
		Namespace: DESCRIPTION_NAMESPACE,
		Vector:    upwork_job_description_vector.Vector,
		TopK:      req.Work_history_count,
		Filter:    filter,
	})
	if err != nil {
		return nil, err
	}
	for _, vector := range response.Matches {
		vector_id := vector.Metadata["freelancer_id"]
		//todo: iterate through all the scores and collate them based off of the freelancer id
		//then we apply weights to the scores based off of fetched data in the next step.
		//then we have apply the same logic for the skill vectors
		if _, exists := scores[vector_id]; !exists {
			scores[vector_id] = []float32{vector.Score}
		} else {
			scores[vector_id] = append(scores[vector_id], vector.Score)
		}
	}

	return scores, nil
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

func (s *UprankVecService) CountTotalWorkHistories(req types.JobEmbeddingData) int {
	total_workhistories := 0
	for _, freelancer := range req.Upwork_job.Edges.UpworkFreelancer {
		total_workhistories += len(freelancer.Edges.WorkHistories)
	}
	return total_workhistories
}

func CreateMetadata(user_id string, job_id string, vector_type string, platform string, freelancer_id *string) map[string]string {
	metadata := make(map[string]string)
	metadata["user_id"] = user_id
	metadata["job_id"] = job_id
	metadata["type"] = vector_type
	metadata["platform"] = platform
	if freelancer_id != nil {
		metadata["freelancer_id"] = *freelancer_id
	}
	return metadata
}
