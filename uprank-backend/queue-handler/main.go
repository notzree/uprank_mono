package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	client "github.com/notzree/uprank-backend/queue-handler/grpc_client"
	"github.com/notzree/uprank-backend/queue-handler/queue"
	"github.com/notzree/uprank-backend/queue-handler/server"
	"github.com/notzree/uprank-backend/queue-handler/service"
)

func main() {
	ranking_queue_url := os.Getenv("RANKING_QUEUE_URL")
	main_backend_url := os.Getenv("MAIN_BACKEND_URL")
	inference_server_url := os.Getenv("INFERENCE_SERVER_URL")
	ms_api_key := os.Getenv("MS_API_KEY")
	queue := queue.NewSqsQueue(ranking_queue_url)
	grpc_inference_client, err := client.NewGRPCInferenceClient(inference_server_url)
	svc := service.NewUprankVecService(main_backend_url, ms_api_key, grpc_inference_client, http.Client{})
	if err != nil {
		log.Fatal(fmt.Sprintf("Error creating grpc client: %v", err))
	}

	server := server.NewServer(queue, svc)
	server.Start()
}
