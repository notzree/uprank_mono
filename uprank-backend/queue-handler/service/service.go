package service

import (
	"context"

	"github.com/notzree/uprank-backend/queue-handler/types"
)

// TODO: add context to all of these so we can cancel the request if needed
type Servicer interface {
	// Upsert vector takes in the job embedding data and then will upsert them into a database. It is expected that Upsert vector only gets passed
	// job embedding data that is not already in the database, or job embedding data that has been updated.
	UpsertVectors(ctx context.Context, req types.JobData, user_id string) (*types.UpsertVectorResponse, error)
	FetchJobData(ctx context.Context, req types.UpworkRankingMessage) (*types.JobData, []types.FreelancerRankingData, error)
	ComputeRawSpecializationScore(ctx context.Context, req types.ComputeRawSpecializationScoreRequest) (*types.ComputeRawSpecializationScoreResponse, error)
	ApplySpecializationScoreWeights(req types.ApplySpecializationScoreWeightsRequest, ctx context.Context, weights ...DescriptionWeight) (*types.ApplySpecializationScoreWeightsResponse, error)
	PostJobRankingData(req types.PostJobRankingDataRequest, ctx context.Context) error
	SaveRawSpecializationScoreWeights(ctx context.Context, req *types.ComputeRawSpecializationScoreResponse, data []types.FreelancerRankingData) error
	SaveWeightedSpecializationScoreWeights(ctx context.Context, req *types.ApplySpecializationScoreWeightsResponse, data []types.FreelancerRankingData) error
	// ComputeEstDuration(ctx context.Context) error
}
