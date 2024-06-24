package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/notzree/uprank-backend/inference-backend/ent"
	"github.com/notzree/uprank-backend/inference-backend/server"
	"github.com/notzree/uprank-backend/inference-backend/service"
	"github.com/pinecone-io/go-pinecone/pinecone"
)

func main() {
	ranking_queue_url := os.Getenv("RANKING_QUEUE_URL")
	db_connection_string := os.Getenv("DB_CONNECTION_STRING")
	pinecone_api_key := os.Getenv("PINECONE_API_KEY")
	sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"))
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}

	ent_client, err := ent.Open("postgres", db_connection_string)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer ent_client.Close()

	pc, pc_err := pinecone.NewClient(pinecone.NewClientParams{
		ApiKey: pinecone_api_key,
	})
	if pc_err != nil {
		log.Fatalf("failed opening connection to pinecone: %v", pc_err)
	}

	sqs_client := sqs.NewFromConfig(sdkConfig)
	queue := server.NewSqsQueue(sqs_client, ranking_queue_url)
	svc := service.NewUpworkService(ent_client)
	server := server.NewServer(ranking_queue_url, queue, svc)
	server.Start()
}
