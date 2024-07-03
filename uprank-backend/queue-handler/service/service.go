package service

import (
	"context"

	"github.com/notzree/uprank-backend/queue-handler/types"
)

// TODO: add context to all of these so we can cancel the request if needed
type Servicer interface {
	// Upsert vector takes in the job embedding data and then will upsert them into a database. It is expected that Upsert vector only gets passed
	// job embedding data that is not already in the database, or job embedding data that has been updated.
	UpsertVectors(req types.JobEmbeddingData, user_id string) (*types.UpsertVectorResponse, error)
	FetchJobData(req types.UpworkRankingMessage) (*types.JobEmbeddingData, error)
	ComputeRawSpecializationScore(req types.ComputeRawSpecializationScoreRequest, ctx context.Context) (*types.ComputeRawSpecializationScoreResponse, error)
	ApplySpecializationScoreWeights(req types.ApplySpecializationScoreWeightsRequest, ctx context.Context, weights ...DescriptionWeight) (*types.ApplySpecializationScoreWeightsResponse, error)
	// ComputeEstDuration(ctx context.Context) error
}
