package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
	proto "github.com/notzree/uprank-backend/queue-handler/proto"
	"github.com/notzree/uprank-backend/queue-handler/queue"
	"github.com/notzree/uprank-backend/queue-handler/types"
)

const (
	DESCRIPTION_NAMESPACE = "description"
	SKILL_NAMESPACE       = "skill"
)

type Server struct {
	messenger_url string
	backend_url   string
	ms_api_key    string
	queue         queue.Queue
	infer         proto.InferenceClient
	httpClient    http.Client
}

func NewServer(messenger_url string, backend_url string, queue queue.Queue, infer proto.InferenceClient, ms_api_key string, httpClient http.Client) *Server {
	return &Server{
		messenger_url: messenger_url,
		backend_url:   backend_url,
		queue:         queue,
		infer:         infer,
		ms_api_key:    ms_api_key,
		httpClient:    httpClient,
	}
}

func (s *Server) Start() {
	log.Println("Starting queue_handler...")
	for {
		Make(s.PollForRankingRequest)()
	}
}

func (s *Server) PollForRankingRequest() error {
	requests, err := s.queue.PollForRankingRequest(context.TODO())
	if err != nil {
		return NewQError(err)
	}
	for _, req := range requests {
		log.Println("Received request:")
		fetched_job_data, err := s.FetchJobData(req)
		if err != nil {
			return NewServiceError(err)
		}
		_, ranking_err := s.UpsertVectors(*fetched_job_data, req.User_id)
		if ranking_err != nil {
			return NewServiceError(ranking_err)
		}
		log.Default().Println("Successfully upserted vectors")

		err = s.queue.DeleteMessage(context.TODO(), req.Receipt_handle)
		if err != nil {
			return NewQError(err)
		}
	}
	return nil
}

func (s *Server) Rank() {
	//todo: impl rank
}

func (s *Server) UpsertVectors(req types.JobDataAll, user_id string) (*types.UpsertVectorResponse, error) {
	ctx := context.Background()
	job_description_vector, embed_err := s.infer.EmbedText(ctx, &proto.EmbedTextRequest{
		Id:       req.ID,
		Text:     req.Description,
		Metadata: CreateMetadata(user_id, req.ID, "job_description", "upwork"),
	})
	if embed_err != nil {
		return nil, embed_err
	}
	job_skills_as_string := strings.Join(req.Skills, " ")
	job_skill_vector, embed_err := s.infer.EmbedText(ctx, &proto.EmbedTextRequest{
		Text:     job_skills_as_string,
		Id:       req.ID,
		Metadata: CreateMetadata(user_id, req.ID, "job_skills", "upwork"),
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

	for _, freelancer := range req.Edges.UpworkFreelancer {
		description_vectors := []*proto.Vector{}
		skill_vectors := []*proto.Vector{}

		//Embed and add vectors to array to embed in 1 trip
		freelancer_description_vector, embed_description_vector_err := s.infer.EmbedText(ctx, &proto.EmbedTextRequest{
			Id:       freelancer.ID,
			Text:     freelancer.Description,
			Metadata: CreateMetadata(user_id, req.ID, "freelancer_description", "upwork"),
		})
		if embed_description_vector_err != nil {
			return nil, embed_description_vector_err
		}

		description_vectors = append(description_vectors, freelancer_description_vector.Vector)

		freelancer_skills_as_string := strings.Join(freelancer.Skills, " ")
		freelancer_skill_vector, embed_skill_vector_err := s.infer.EmbedText(ctx, &proto.EmbedTextRequest{
			Id:       freelancer.ID,
			Text:     freelancer_skills_as_string,
			Metadata: CreateMetadata(user_id, req.ID, "freelancer_skills", "upwork"),
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
				Metadata: CreateMetadata(user_id, req.ID, "work_history_description", "upwork"),
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
	}

	return &types.UpsertVectorResponse{
		Job_description_vector: job_description_vector.Vector,
		Job_skill_vector:       job_skill_vector.Vector,
	}, nil

}

func (s *Server) FetchJobData(req types.UpworkRankingMessage) (*types.JobDataAll, error) {
	fetch_url := fmt.Sprintf("%s/v1/private/jobs/%s/%s/all_data", s.backend_url, req.Platform, req.Platform_id)
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
	var job_data types.JobDataAll
	err = json.Unmarshal((body), &job_data)
	if err != nil {
		return nil, err
	}
	return &job_data, nil
}

func CreateMetadata(user_id string, job_id string, vector_type string, platform string) map[string]string {
	metadata := make(map[string]string)
	metadata["user_id"] = user_id
	metadata["job_id"] = job_id
	metadata["type"] = vector_type
	metadata["platform"] = platform
	return metadata
}
