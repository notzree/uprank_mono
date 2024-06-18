package service

import (
	"context"
	"log"

	"github.com/notzree/uprank-backend/inference-backend/ent"
	"github.com/notzree/uprank-backend/inference-backend/ent/job"
	"github.com/notzree/uprank-backend/inference-backend/ent/user"

	"github.com/notzree/uprank-backend/inference-backend/types"
)

type UpworkService struct {
	ent *ent.Client
}

func NewUpworkService(ent *ent.Client) *UpworkService {
	return &UpworkService{ent: ent}
}

func (u *UpworkService) RankUpworkJob(data types.UpworkRankingMessage, ctx context.Context) (*types.UpworkRankingResult, error) {
	job_data, err := u.ent.Job.Query().
		Where(job.IDEQ(data.Job_id)).
		Where(job.HasUserWith(user.IDEQ(data.User_id))).
		WithUpworkjob().
		QueryUpworkjob().WithUpworkfreelancer().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	log.Default().Printf("%+v\n", job_data.Edges.Upworkfreelancer)

	return nil, nil
}
