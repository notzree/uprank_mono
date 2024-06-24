package service

import (
	"context"

	"github.com/notzree/uprank-backend/inference-backend/types"
)

type Service interface {
	RankUpworkJob(data types.UpworkRankingMessage, ctx context.Context) (*types.UpworkRankingResult, error)
}
