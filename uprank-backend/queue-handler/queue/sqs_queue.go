package queue

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/google/uuid"
	types "github.com/notzree/uprank-backend/queue-handler/types"
)

type SqsQueue struct {
	sqs_client        *sqs.Client
	ranking_queue_url string
}

func NewSqsQueue(ranking_queue_url string) *SqsQueue {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"))
	if err != nil {
		log.Panic(err)
		log.Fatal("Couldn't load default configuration. Have you set up your AWS account?")
	}

	sqs_client := sqs.NewFromConfig(sdkConfig)
	return &SqsQueue{
		sqs_client:        sqs_client,
		ranking_queue_url: ranking_queue_url,
	}
}

func (s *SqsQueue) PollForRankingRequest(ctx context.Context) ([]types.UpworkRankingMessage, error) {
	ranking_messages := []types.UpworkRankingMessage{}
	response, err := s.sqs_client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl: &s.ranking_queue_url,
		MessageAttributeNames: []string{
			"job_id",
			"user_id",
			"platform",
			"platform_id",
		},
		WaitTimeSeconds: 20,
	})
	if err != nil {
		return nil, err
	}

	for _, message := range response.Messages {
		ranking_messages = append(ranking_messages, types.UpworkRankingMessage{
			Job_id:         uuid.Must(uuid.Parse(*message.MessageAttributes["job_id"].StringValue)),
			User_id:        *message.MessageAttributes["user_id"].StringValue,
			Platform:       *message.MessageAttributes["platform"].StringValue,
			Platform_id:    *message.MessageAttributes["platform_id"].StringValue,
			Message:        *message.Body,
			Receipt_handle: *message.ReceiptHandle,
		})

	}
	return ranking_messages, nil
}

func (s *SqsQueue) SendRankingResult(ctx context.Context) error {
	// UNIMPLEMENTED
	return nil
}

func (s *SqsQueue) DeleteMessage(ctx context.Context, receipt_handle string) error {
	_, err := s.sqs_client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      &s.ranking_queue_url,
		ReceiptHandle: &receipt_handle,
	})
	if err != nil {
		return err
	}
	return nil
}
