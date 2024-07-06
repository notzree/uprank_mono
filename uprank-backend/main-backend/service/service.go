package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/notzree/uprank-backend/main-backend/ent"
	"github.com/notzree/uprank-backend/main-backend/types"
)

type Servicer interface {
	CreateUser(data types.CreateUserRequest, ctx context.Context) (*ent.User, error)
	UpdateUser(data types.UpdateUserRequest, ctx context.Context) (*ent.User, error)
	CreateJob(data types.CreateJobRequest, user_id string, ctx context.Context) (*ent.Job, error)
	GetJobs(user_id string, ctx context.Context) ([]*ent.Job, error)
	AttachPlatformSpecificjobs(data types.AttachPlatformSpecificJobsRequest, user_id string, job_id *uuid.UUID, ctx context.Context) (*ent.Job, error)
	GetUpworkJob(upwork_job_id string, user_id string, ctx context.Context) (*ent.UpworkJob, error)
	GetUpworkJobEmbeddingData(upwork_job_id string, user_id string, ctx context.Context) (*ent.UpworkJob, error)
	UpdateUpworkJob(data types.UpdateUpworkJobRequest, user_id string, ctx context.Context) (*ent.UpworkJob, error)
	CreateUpworkFreelancer(data []types.CreateUpworkFreelancerRequest, user_id string, upwork_job_id string, ctx context.Context) ([]*ent.UpworkFreelancer, error)
	UpsertUpworkFreelancer(data []types.CreateUpworkFreelancerRequest, user_id string, upwork_job_id string, ctx context.Context) (created_freelancer_count int, updated_freelancer_count int, deleted_freelancer_count int, err error)
	UpdateUpworkFreelancer(data []types.UpdateUpworkFreelancerRequest, user_id, upwork_job_id string, ctx context.Context) ([]string, error)
	GetFreelancersFromUpworkJob(upwork_job_id string, user_id string, ctx context.Context) ([]*ent.UpworkFreelancer, error)
	SendRankingrequest(data types.RankJobRequest, ctx context.Context) error //sends request to messaging queue to rank the job
	AddJobRankings(data types.AddJobRankingRequest, user_id string, ctx context.Context) error
}
