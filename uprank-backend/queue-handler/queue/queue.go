package queue

import (
	"context"

	types "github.com/notzree/uprank-backend/queue-handler/types"
)

type Queue interface {
	PollForRankingRequest(ctx context.Context) ([]types.UpworkRankingMessage, error)
	DeleteMessage(ctx context.Context, receipt_handle string) error
	SendRankingResult(ctx context.Context) error
}