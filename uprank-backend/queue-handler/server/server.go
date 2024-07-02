package server

import (
	"context"
	"log"
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
		compute_specialization_req := types.ComputeSpecializationRequest{
			Job_id:                 req.Job_id.String(),
			Work_history_count:     int32(total_work_histories),
			Job_description_vector: upsert_resp.Job_description_vector,
			Job_skill_vector:       upsert_resp.Job_skill_vector,
		}
		_, compute_specialization_err := s.svc.ComputeSpecialization(compute_specialization_req, context.TODO())
		if compute_specialization_err != nil {
			return NewServiceError(compute_specialization_err)
		}

		err = s.queue.DeleteMessage(context.TODO(), req.Receipt_handle)
		if err != nil {
			return NewQError(err)
		}
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
