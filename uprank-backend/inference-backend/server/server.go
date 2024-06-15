package server

import (
	"context"
	"log"
)

type Server struct {
	messenger_url string
	queue         Queue
}

func NewServer(messenger_url string, queue Queue) *Server {
	return &Server{
		messenger_url: messenger_url,
		queue:         queue,
	}
}

func (s *Server) Start() error {
	//So in the future we can add more listners
	return s.ListenForRankingRequests()
}

func (s *Server) ListenForRankingRequests() error {
	log.Println("Listening for ranking requests")
	for {
		requests, err := s.queue.ReceiveRankingRequest(context.TODO())
		if err != nil {
			return err
		}
		for _, req := range requests {
			log.Default().Printf("Received message: %s", req)
			err := s.queue.DeleteMessage(context.TODO(), req.Receipt_handle)
			if err != nil {
				return err
			}
		}

	}
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
