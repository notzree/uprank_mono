package service

import (
	"context"

	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/job"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/upworkjob"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/user"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/types"
)

type UpworkJobService struct {
	ent *ent.Client
}

type NewUpworkJobServiceParams struct {
	Ent *ent.Client
}

func NewUpworkJobService(params NewUpworkJobServiceParams) UpworkJobService {
	return UpworkJobService{
		ent: params.Ent,
	}
}

func (s *UpworkJobService) GetAllUpworkJobsForUser(ctx context.Context, params types.QueryUpworkJobRequest) (*ent.UpworkJob, error) {
	job, err := s.ent.UpworkJob.Query().
		Where(upworkjob.IDEQ(params.Upwork_job_id)).
		Where(
			upworkjob.HasJobWith(
				job.HasUserWith(user.IDEQ(params.User_id)),
			),
		).
		WithUpworkfreelancer().
		Only(ctx)
	return job, err
}

func (s *UpworkJobService) GetUpworkJobEmbeddingData(ctx context.Context, params types.QueryUpworkJobRequest) (*ent.UpworkJob, error) {
	job, err := s.ent.UpworkJob.Query().
		Where(upworkjob.IDEQ(params.Upwork_job_id)).
		Where(
			upworkjob.HasJobWith(
				job.HasUserWith(user.IDEQ(params.User_id)),
			),
		).
		WithUpworkfreelancer(func(query *ent.UpworkFreelancerQuery) {
			query.WithAttachments()
			query.WithWorkHistories()
		}).
		Only(ctx)
	return job, err
}

func (s *UpworkJobService) UpdateUpworkJob(ctx context.Context, params types.UpdateUpworkJobRequest) (*ent.UpworkJob, error) {
	upwork_job_data := params.Data
	placeholder_upwork_job := s.ent.UpworkJob.UpdateOneID(upwork_job_data.Upwork_id).
		Where(upworkjob.HasJobWith(job.HasUserWith(user.IDEQ(params.User_id))))
	if upwork_job_data.Title != nil {
		placeholder_upwork_job.SetTitle(*upwork_job_data.Title)
	}
	if upwork_job_data.Location != nil {
		placeholder_upwork_job.SetLocation(*upwork_job_data.Location)
	}
	if upwork_job_data.Description != nil {
		placeholder_upwork_job.SetDescription(*upwork_job_data.Description)
	}
	if upwork_job_data.Skills != nil {
		placeholder_upwork_job.SetSkills(*upwork_job_data.Skills)
	}
	if upwork_job_data.Embedded_at != nil {
		placeholder_upwork_job.SetEmbeddedAt(*upwork_job_data.Embedded_at)
	}
	if upwork_job_data.Ranked_at != nil {
		placeholder_upwork_job.SetRankedAt(*upwork_job_data.Ranked_at)
	}
	if upwork_job_data.Experience_level != nil {
		placeholder_upwork_job.SetExperienceLevel(*upwork_job_data.Experience_level)
	}
	updated_upwork_job, err := placeholder_upwork_job.Save(ctx)

	return updated_upwork_job, err
}
