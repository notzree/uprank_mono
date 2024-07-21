package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/job"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/schema"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/user"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/types"
)

type JobService struct {
	ent *ent.Client
}

type NewJobServiceParams struct {
	Ent *ent.Client
}

func NewJobService(params NewJobServiceParams) JobService {
	return JobService{
		ent: params.Ent,
	}
}

// Creates a job and for each non-null platform job request, creates a platform job and attaches it to the job.
func (s *JobService) CreateJob(ctx context.Context, params types.CreateJobRequest) (*ent.Job, error) {
	new_job, create_job_err := s.ent.Job.Create().SetUserID(params.User_id).SetOriginPlatform(schema.Platform(params.Origin)).Save(ctx) //TODO: Make the origin platform dynamic
	if create_job_err != nil {
		return nil, create_job_err
	}

	_, attach_platform_jobs_err := s.AttachPlatformSpecificjobs(ctx, types.AttachPlatformSpecificJobsRequest{
		User_id:                 params.User_id,
		Job_id:                  new_job.ID,
		PlatformSpecificJobdata: params.PlatformJobRequests,
	})
	if attach_platform_jobs_err != nil {
		return nil, attach_platform_jobs_err
	}
	return new_job, nil
}

func (s *JobService) AttachPlatformSpecificjobs(ctx context.Context, params types.AttachPlatformSpecificJobsRequest) (*ent.Job, error) {
	if params.PlatformSpecificJobdata.UpworkRequest != nil {
		_, create_upwork_job_err := s.AttachUpworkJob(ctx, types.AttachUpworkJobRequest{
			Job_id:          params.Job_id,
			User_id:         params.User_id,
			Upwork_job_data: *params.PlatformSpecificJobdata.UpworkRequest,
		})
		if create_upwork_job_err != nil {
			return nil, create_upwork_job_err
		}
	}
	//add more platforms here
	return nil, nil
}

func (s *JobService) AttachUpworkJob(ctx context.Context, params types.AttachUpworkJobRequest) (*ent.UpworkJob, error) {
	upwork_job := &params.Upwork_job_data
	new_upwork_job, create_upwork_job_err := s.ent.UpworkJob.Create().
		SetID(upwork_job.Upwork_job_id).
		SetTitle(upwork_job.Title).
		SetLocation(upwork_job.Location).
		SetDescription(upwork_job.Description).
		SetSkills(upwork_job.Skills).
		SetExperienceLevel(upwork_job.Experience_level).
		SetHourly(upwork_job.Hourly).
		SetFixed(upwork_job.Fixed).
		SetHourlyRate(upwork_job.Hourly_rate).
		SetFixedRate(upwork_job.Fixed_rate).
		SetJobID(params.Job_id).
		AddUserIDs(params.User_id).
		Save(ctx)
	if create_upwork_job_err != nil {
		return nil, create_upwork_job_err
	}
	return new_upwork_job, create_upwork_job_err
}

func (s *JobService) GetAllJobsForUser(ctx context.Context, params types.GetAllJobsForUserRequest) ([]*ent.Job, error) {
	jobs, err := s.ent.Job.Query().Where(job.HasUserWith(user.IDEQ(params.User_id))).WithUpworkjob().All(ctx)
	return jobs, err
}

func (s *JobService) GetJobById(ctx context.Context, params types.GetJobByIdRequest) (*ent.Job, error) {
	job, err := s.ent.Job.Query().Where(job.IDEQ(uuid.MustParse(params.Job_id))).Where(job.HasUserWith(user.IDEQ(params.User_id))).WithUpworkjob(func(query *ent.UpworkJobQuery) {
		query.WithUpworkfreelancer(func(query *ent.UpworkFreelancerQuery) {
			query.WithFreelancerInferenceData()
		})
	}).Only(ctx)
	return job, err
}
