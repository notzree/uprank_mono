package server

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/notzree/uprank-backend/queue-handler/queue"
	"github.com/notzree/uprank-backend/queue-handler/service"
	"github.com/notzree/uprank-backend/queue-handler/types"
)

type Server struct {
	queue queue.Queue
	svc   service.Servicer
}

func NewServer(queue queue.Queue, svc service.Servicer) *Server {
	return &Server{
		queue: queue,
		svc:   svc,
	}
}

// in the future if we have more than 1 queue then we should delegate each listener to a go routine
func (s *Server) Start() {
	log.Println("Starting queue_handler...")
	for {
		err := s.PollForRankingRequest()
		if err != nil {
			HandleError(err)
		}
	}
}

func (s *Server) PollForRankingRequest() error {
	requests, err := s.queue.PollForRankingRequest(context.TODO())
	if err != nil {
		return NewQError(err)
	}
	for _, req := range requests {
		//TODO: MAKE THIS A GO ROUTINE
		log.Println("Received request")
		ctx := context.TODO()
		go func(ctx context.Context, req types.UpworkRankingMessage) {
			err := s.HandleRankingRequest(ctx, req)
			if err != nil {
				HandleError(err)
			}

		}(ctx, req)
	}
	return nil
}

func (s *Server) HandleRankingRequest(ctx context.Context, req types.UpworkRankingMessage) error {
	fetched_job_data, err := s.svc.FetchJobData(req)
	if err != nil {
		return NewServiceError(err)
	}
	start := time.Now()
	upsert_resp, upsert_err := s.svc.UpsertVectors(*fetched_job_data, req.User_id)
	if upsert_err != nil {
		return NewServiceError(upsert_err)
	}
	elapsed := time.Since(start)
	log.Printf("Upsert time: %s", elapsed)

	total_work_histories := s.CountTotalWorkHistories(*fetched_job_data)
	compute_specialization_req := types.ComputeRawSpecializationScoreRequest{
		Job_id:                 req.Job_id.String(),
		Work_history_count:     int32(total_work_histories),
		Freelancer_count:       int32(len(fetched_job_data.Upwork_job.Edges.UpworkFreelancer)),
		Job_description_vector: upsert_resp.Job_description_vector,
	}

	raw_specialization_scores, compute_raw_specialization_err := s.svc.ComputeRawSpecializationScore(compute_specialization_req, context.TODO())
	if compute_raw_specialization_err != nil {
		return NewServiceError(compute_raw_specialization_err)
	}

	final_specialization_scores, apply_specialization_weights_err := s.svc.ApplySpecializationScoreWeights(types.ApplySpecializationScoreWeightsRequest{
		Description_scores: *raw_specialization_scores.Job_description_specialization_scores,
		Job_data:           *fetched_job_data,
	}, context.TODO())
	if apply_specialization_weights_err != nil {
		return NewServiceError(apply_specialization_weights_err)
	}
	// Write final_specialization_scores to a file as JSON
	file, err := os.Create("final_specialization_scores.json")
	if err != nil {
		return NewServiceError(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(final_specialization_scores)
	if err != nil {
		return NewServiceError(err)
	}

	err = s.queue.DeleteMessage(context.TODO(), req.Receipt_handle)
	if err != nil {
		return NewQError(err)
	}
	return nil
}

func (s *Server) CountTotalWorkHistories(req types.JobEmbeddingData) int {
	total_workhistories := 0
	for _, freelancer := range req.Upwork_job.Edges.UpworkFreelancer {
		total_workhistories += len(freelancer.Edges.WorkHistories)
	}
	return total_workhistories

}
