package main

import (
	"fmt"
	"log"
	"os"

	client "github.com/notzree/uprank-backend/queue-handler/grpc_client"
	"github.com/notzree/uprank-backend/queue-handler/queue"
	"github.com/notzree/uprank-backend/queue-handler/server"
)

func main() {
	ranking_queue_url := os.Getenv("RANKING_QUEUE_URL")
	grpc_server_url := ""
	queue := queue.NewSqsQueue(ranking_queue_url)
	grpc_inference_client, err := client.NewGRPCInferenceClient(grpc_server_url)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error creating grpc client: %v", err))
	}

	server := server.NewServer(ranking_queue_url, queue, grpc_inference_client)
	server.Start()
}

//server
//server -> queue, service
//queue -> handles polling for messages
//services -> business logic

// service -> grpc_client -> interact with python server

//server
//grpc_client
//queue
