package server

import (
	"context"

	types "github.com/notzree/uprank-backend/inference-backend/types"
)

type Queue interface {
	ReceiveRankingRequest(ctx context.Context) ([]types.UpworkRankingMessage, error)
	DeleteMessage(ctx context.Context, receipt_handle string) error
	SendRankingResult(ctx context.Context) error
}
