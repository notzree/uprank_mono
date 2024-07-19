package server

import (
	"context"
	"log"

	_ "github.com/lib/pq"
	"github.com/notzree/uprank_mono/uprank-backend/queue-handler/queue"
	"github.com/notzree/uprank_mono/uprank-backend/queue-handler/service"
	"github.com/notzree/uprank_mono/uprank-backend/queue-handler/types"
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
