package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/notzree/uprank-backend/inference-backend/server"
)

func main() {
	ranking_queue_url := os.Getenv("RANKING_QUEUE_URL")
	sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"))
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}
	sqs_client := sqs.NewFromConfig(sdkConfig)
	queue := server.NewSqsQueue(sqs_client, ranking_queue_url)
	server := server.NewServer(ranking_queue_url, queue)
	log.Fatal(server.Start())
}
