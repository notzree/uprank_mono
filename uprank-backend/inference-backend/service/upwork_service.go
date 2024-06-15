package service

import (
	"context"

	"github.com/notzree/uprank-backend/inference-backend/ent"
	"github.com/notzree/uprank-backend/inference-backend/types"
)

type UpworkService struct {
	ent *ent.Client
}

func (u *UpworkService) RankUpworkJob(data types.UpworkRankingMessage, ctx context.Context) (types.UpworkRankingResult, error) {

}
