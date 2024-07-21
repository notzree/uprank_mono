package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/types"
)

func (s *Server) AttachPlatformSpecificJobs(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	user_id, user_id_err := s.authenticator.GetIdFromRequest(r)
	if user_id_err != nil {
		return user_id_err
	}
	jobId := chi.URLParam(r, "job_id")
	jobIdUUID, err := uuid.Parse(jobId)
	if err != nil {
		return InvalidJSON()
	}
	var req types.AttachPlatformSpecificJobsData

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}

	defer r.Body.Close()
	if errors := req.Validate(); len(errors) > 0 {
		return InvalidRequestData(errors)
	}
	new_upwork_job, err := s.job.AttachPlatformSpecificjobs(ctx, types.AttachPlatformSpecificJobsRequest{
		User_id:                 user_id,
		Job_id:                  jobIdUUID,
		PlatformSpecificJobdata: req,
	})
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, new_upwork_job)
}

func (s *Server) CreateJob(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	user_id, user_id_err := s.authenticator.GetIdFromRequest(r)
	if user_id_err != nil {
		return user_id_err
	}
	var req types.CreateJobRequest // Uhh this shouldn't have user id?? but ok for now?

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}

	if errors := req.Validate(); len(errors) > 0 {
		return InvalidRequestData(errors)
	}

	if req.User_id != user_id {
		return InvalidJSON()
	}
	new_job, err := s.job.CreateJob(ctx, req)

	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, new_job)
}

func (s *Server) GetJobs(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	user_id, user_id_err := s.authenticator.GetIdFromRequest(r)
	if user_id_err != nil {
		return user_id_err
	}

	jobs, err := s.job.GetAllJobsForUser(ctx, types.GetAllJobsForUserRequest{
		User_id: user_id,
	})
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, jobs)
}

func (s *Server) GetJobById(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	user_id, user_id_err := s.authenticator.GetIdFromRequest(r)
	if user_id_err != nil {
		return user_id_err
	}
	jobId := chi.URLParam(r, "job_id")
	job, err := s.job.GetJobById(ctx, types.GetJobByIdRequest{
		Job_id:  jobId,
		User_id: user_id,
	})
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, job)
}
