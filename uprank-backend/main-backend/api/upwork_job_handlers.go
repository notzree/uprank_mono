package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/types"
)

func (s *Server) GetUpworkJob(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	userId, user_id_err := s.authenticator.GetIdFromRequest(r)
	if user_id_err != nil {
		return user_id_err
	}
	upworkJobId := chi.URLParam(r, "upwork_job_id")
	job, err := s.upwork_job.GetAllUpworkJobsForUser(ctx, types.QueryUpworkJobRequest{
		Upwork_job_id: upworkJobId,
		User_id:       userId,
	})
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, job)
}

func (s *Server) GetUpworkJobEmbeddingData(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	user_id, user_id_err := s.authenticator.GetIdFromRequest(r)
	if user_id_err != nil {
		return user_id_err
	}
	upwork_job_id := chi.URLParam(r, "upwork_job_id")
	job_id := chi.URLParam(r, "job_id")

	job, err := s.upwork_job.GetUpworkJobEmbeddingData(ctx, types.QueryUpworkJobRequest{
		Upwork_job_id: upwork_job_id,
		User_id:       user_id,
	})
	if err != nil {
		return err
	}
	JobEmbeddingData := &types.GetUpworkJobEmbeddingDataResponse{
		Job_id:     job_id,
		Upwork_job: job,
	}
	return writeJSON(w, http.StatusOK, JobEmbeddingData)
}

func (s *Server) UpdateUpworkJob(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	user_id, user_id_err := s.authenticator.GetIdFromRequest(r)
	if user_id_err != nil {
		return user_id_err
	}
	var req types.UpdateUpworkJobData

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}

	defer r.Body.Close()
	if errors := req.Validate(); len(errors) > 0 {
		return InvalidRequestData(errors)
	}
	upwork_job, err := s.upwork_job.UpdateUpworkJob(ctx, types.UpdateUpworkJobRequest{
		User_id: user_id,
		Data:    req,
	})
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, upwork_job)
}
