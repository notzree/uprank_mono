package service

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	sqs_types "github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/types"
)

type RankingService struct {
	ent               *ent.Client
	sqs_client        *sqs.Client
	ranking_queue_url string
}

type NewRankingServiceParams struct {
	Ent               *ent.Client
	Sqs_client        *sqs.Client
	Ranking_queue_url string
}

func NewRankingService(params NewRankingServiceParams) RankingService {
	return RankingService{
		ent:               params.Ent,
		sqs_client:        params.Sqs_client,
		ranking_queue_url: params.Ranking_queue_url,
	}
}

func (s *RankingService) AddJobRankings(ctx context.Context, params types.AddJobRankingRequest) error {
	bulk := make([]*ent.FreelancerInferenceDataCreate, 0, len(params.Freelancer_ranking_data))
	for _, inference_data := range params.Freelancer_ranking_data {
		bulk = append(bulk, s.ent.FreelancerInferenceData.Create().
			SetUpworkfreelancerID(inference_data.Freelancer_id).
			SetFinalizedRatingScore(float64(inference_data.Finalized_rating_score)).
			SetUprankReccomended(inference_data.Uprank_reccomended).
			SetUprankNotEnoughData(inference_data.Uprank_not_enough_data).
			SetUprankReccomendedReasons(inference_data.Uprank_reccomended_reasons).
			SetRawRatingScore(float64(inference_data.Raw_rating_score)).
			SetBudgetAdherencePercentage(float64(inference_data.Budget_adherence_percentage)).
			SetBudgetOverrunPercentage(float64(inference_data.Budget_overrun_percentage)))
	}
	err := s.ent.FreelancerInferenceData.CreateBulk(bulk...).OnConflict().DoNothing().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

// So Basically, the fifo queue uses the messagebody as the deduplication id so if we send a message with the same body within a span of 5 mins it will be discarded.
func (s *RankingService) SendRankingrequest(ctx context.Context, data types.RankJobRequest) error {
	_, err := s.sqs_client.SendMessage(ctx, &sqs.SendMessageInput{
		MessageAttributes: map[string]sqs_types.MessageAttributeValue{
			"job_id": {
				DataType:    aws.String("String"),
				StringValue: aws.String(data.Job_id.String()),
			},
			"user_id": {
				DataType:    aws.String("String"),
				StringValue: aws.String(data.User_id),
			},
			"short_lived_token": {
				DataType:    aws.String("String"),
				StringValue: aws.String(data.Short_lived_token),
			},
			"platform": {
				DataType:    aws.String("String"),
				StringValue: aws.String(data.Platform),
			},
			"platform_id": {
				DataType:    aws.String("String"),
				StringValue: aws.String(data.Platform_id),
			},
		},
		QueueUrl:               &s.ranking_queue_url,
		MessageBody:            aws.String(fmt.Sprint("Ranking request for job ", data.Job_id, " by user ", data.User_id)),
		MessageDeduplicationId: aws.String(data.Job_id.String()),
		MessageGroupId:         aws.String("RankingRequest"),
	})
	if err != nil {
		return err
	}

	return nil
}
