package server

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/google/uuid"
	types "github.com/notzree/uprank-backend/inference-backend/types"
)

type SqsQueue struct {
	sqs_client *sqs.Client
	queue_url  string
}

func NewSqsQueue(sqs_client *sqs.Client, queue_url string) *SqsQueue {
	return &SqsQueue{
		sqs_client: sqs_client,
		queue_url:  queue_url,
	}
}

func (s *SqsQueue) ReceiveRankingRequest(ctx context.Context) ([]types.UpworkRankingMessage, error) {
	ranking_messages := []types.UpworkRankingMessage{}
	response, err := s.sqs_client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl: &s.queue_url,
		MessageAttributeNames: []string{
			"job_id",
			"user_id",
		},
		WaitTimeSeconds: 5,
	})
	if err != nil {
		return nil, err
	}

	for _, message := range response.Messages {
		ranking_messages = append(ranking_messages, types.UpworkRankingMessage{
			Job_id:         uuid.Must(uuid.Parse(*message.MessageAttributes["job_id"].StringValue)),
			User_id:        *message.MessageAttributes["user_id"].StringValue,
			Message:        *message.Body,
			Receipt_handle: *message.ReceiptHandle,
		})

	}
	return ranking_messages, nil
}

func (s *SqsQueue) SendRankingResult(ctx context.Context) error {
	return nil
}

func (s *SqsQueue) DeleteMessage(ctx context.Context, receipt_handle string) error {
	_, err := s.sqs_client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      &s.queue_url,
		ReceiptHandle: &receipt_handle,
	})
	if err != nil {
		return err
	}
	return nil
}
