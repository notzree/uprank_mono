package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
	proto "github.com/notzree/uprank-backend/queue-handler/proto"
	"github.com/notzree/uprank-backend/queue-handler/queue"
	"github.com/notzree/uprank-backend/queue-handler/types"
)

type Server struct {
	messenger_url string
	backend_url   string
	queue         queue.Queue
	infer         proto.InferenceClient
}

func NewServer(messenger_url string, backend_url string, queue queue.Queue, infer proto.InferenceClient) *Server {
	return &Server{
		messenger_url: messenger_url,
		backend_url:   backend_url,
		queue:         queue,
		infer:         infer,
	}
}

func (s *Server) Start() {
	//So in the future we can add more listners
	//also need to figure out middleware and logging
	//maybe have a makefunction here that wraps functions that returns errors into return void
	//wrapping this with make makes it not run
	log.Println("Listening for requests")
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

		err = s.queue.DeleteMessage(context.TODO(), req.Receipt_handle)
		if err != nil {
			return NewQError(err)
		}
	}
	return nil
}

func (s *Server) UpsertVectors(req types.JobDataAll, user_id string) (*string, error) {
	ctx := context.Background()
	job_description_vector, err := s.infer.EmbedText(ctx, &proto.EmbedTextRequest{
		Text: req.Description,
	})
	description_metadata := CreateMetadata(user_id, req.ID, "job_description", "upwork")

	job_description_pc_vector := &proto.Vector{
		Id:       req.ID,
		Vector:   job_description_vector.Vector,
		Metadata: description_metadata,
	}

	description_upsert_result, err := s.infer.UpsertVector(ctx, &proto.UpsertVectorRequest{
		Vectors: []*proto.Vector{job_description_pc_vector},
	})
	if err != nil {
		return nil, err
	}
	job_skills_as_string := strings.Join(req.Skills, " ")
	job_skill_vector, err := s.infer.EmbedText(ctx, &proto.EmbedTextRequest{
		Text: job_skills_as_string,
	})
	if err != nil {
		return nil, err
	}
	skill_metadata := CreateMetadata(user_id, req.ID, "job_skills", "upwork")

	job_skill_pc_vector := &proto.Vector{
		Id:       req.ID,
		Vector:   job_skill_vector.Vector,
		Metadata: skill_metadata,
	}

	skills_upsert_result, err := s.infer.UpsertVector(ctx, &proto.UpsertVectorRequest{
		Vectors: []*proto.Vector{job_skill_pc_vector},
	})
	if err != nil {
		return nil, err
	}

	for _, freelancer := range req.Edges.UpworkFreelancer {
		freelancer_description_vector, embed_description_vector_err := s.infer.EmbedText(ctx, &proto.EmbedTextRequest{
			Text: freelancer.Description,
		})
		if embed_description_vector_err != nil {
			return nil, embed_description_vector_err
		}
		freelancer_skills_as_string := strings.Join(freelancer.Skills, " ")
		freelancer_skill_vector, embed_skill_vector_err := s.infer.EmbedText(ctx, &proto.EmbedTextRequest{
			Text: freelancer_skills_as_string,
		})
		if embed_skill_vector_err != nil {
			return nil, embed_skill_vector_err
		}
		for _, work_history := range freelancer.Edges.WorkHistories {

		}

		//TODO: Implement upserting of the work histories
		//todo: prob should have a batch upsert option
		//Also should have a way to embed + upsert in 1 round trip so better performance

	}

}

func (s *Server) FetchJobData(req types.UpworkRankingMessage) (*types.JobDataAll, error) {
	fetch_url := fmt.Sprintf("%s/v1/private/jobs/%s/%s/all_data", s.backend_url, req.Platform, req.Platform_id)
	log.Println("Fetching data from:", fetch_url)
	resp, err := http.Get(fetch_url)
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

func CreateMetadata(user_id string, job_id string, job_type string, platform string) map[string]string {
	metadata := make(map[string]string)
	metadata["user_id"] = user_id
	metadata["job_id"] = job_id
	metadata["type"] = job_type
	metadata["platform"] = platform
	return metadata
}
