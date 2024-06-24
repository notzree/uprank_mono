package server

import (
	"context"
	"log"

	_ "github.com/lib/pq"
	"github.com/notzree/uprank-backend/inference-backend/service"
)

type Server struct {
	messenger_url string
	queue         Queue
	svc           service.Service
}

func NewServer(messenger_url string, queue Queue, svc service.Service) *Server {
	return &Server{
		messenger_url: messenger_url,
		queue:         queue,
		svc:           svc,
	}
}

func (s *Server) Start() {
	//So in the future we can add more listners
	//also need to figure out middleware and logging
	//maybe have a makefunction here that wraps functions that returns errors into return void
	//wrapping this with make makes it not run
	log.Println("Listening for requests")
	for {
		Make(s.ListenForRankingRequests)()
	}
}

func (s *Server) ListenForRankingRequests() error {
	requests, err := s.queue.ReceiveRankingRequest(context.TODO())
	if err != nil {
		return NewQError(err)
	}
	for _, req := range requests {
		log.Println("Received request:")
		_, ranking_err := s.svc.RankUpworkJob(req, context.TODO())
		if ranking_err != nil {
			return NewServiceError(ranking_err)
		}

		err := s.queue.DeleteMessage(context.TODO(), req.Receipt_handle)
		if err != nil {
			return NewQError(err)
		}
	}
	return nil
}

//thinking
//Want some S.O.C
//the queue class handles all the processing
//the service class handles all the logic
//In the listen for ranking req function:
//Need to call queue.ReceiveRankingRequest
//then take the res of that and call svc.Rank(req)
//then if the result of the rank is sucessful
//queue.DeleteMessage
//then queue.SendRankingResult
//Need to figure out logging!?
