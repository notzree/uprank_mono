package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/notzree/uprank-backend/main-backend/ent"
	"github.com/notzree/uprank-backend/main-backend/types"
)

type PlatformFreelancerCreator func(data types.CreateFreelancersRequest, ctx context.Context) (*ent.UpworkFreelancer, error) //TODO: Figure out how to chagne freelancers and then do this.

type Servicer interface {
	CreateUser(data types.CreateUserRequest, ctx context.Context) (*ent.User, error)
	UpdateUser(data types.UpdateUserRequest, ctx context.Context) (*ent.User, error)
	CreateJob(data types.CreateJobRequest, user_id string, ctx context.Context) (*ent.Job, error)
	AttachUpworkJob(data types.AttachUpworkJobRequest, user_id string, job_id *uuid.UUID, ctx context.Context) (*ent.UpworkJob, error)
	GetUpworkJob(upwork_job_id string, user_id string, ctx context.Context) (*ent.UpworkJob, error)
	// UpdateUpworkJob(data types.CreateUpworkJobRequest, user_id string, ctx context.Context) (*ent.Job, *ent.UpworkJob, error)
	CreateUpworkFreelancer(data []types.CreateUpworkFreelancerRequest, user_id string, upwork_job_id string, ctx context.Context) ([]*ent.UpworkFreelancer, error)
	UpdateUpworkFreelancer(data []types.CreateUpworkFreelancerRequest, user_id string, upwork_job_id string, ctx context.Context) (created_freelancer_count int, updated_freelancer_count int, deleted_freelancer_count int, err error)
}
