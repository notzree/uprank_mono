package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	client "github.com/notzree/uprank-backend/queue-handler/grpc_client"
	"github.com/notzree/uprank-backend/queue-handler/queue"
	"github.com/notzree/uprank-backend/queue-handler/server"
)

func main() {
	ranking_queue_url := os.Getenv("RANKING_QUEUE_URL")
	main_backend_url := os.Getenv("MAIN_BACKEND_URL")
	// inference_server_url := os.Getenv("INFERENCE_SERVER_URL")
	// log.Default().Print(inference_server_url)
	inference_server_url := "uprank-inference-backend:50051"
	ms_api_key := os.Getenv("MS_API_KEY")
	queue := queue.NewSqsQueue(ranking_queue_url)
	grpc_inference_client, err := client.NewGRPCInferenceClient(inference_server_url)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error creating grpc client: %v", err))
	}

	client := http.Client{}
	server := server.NewServer(ranking_queue_url, main_backend_url, queue, grpc_inference_client, ms_api_key, client)
	server.Start()
}
