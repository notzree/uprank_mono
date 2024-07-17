package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/types"
)

func (s *Server) AttachPlatformSpecificJobs(w http.ResponseWriter, r *http.Request) error {
	user_id, user_id_err := s.authenticator.GetIdFromRequest(r)
	if user_id_err != nil {
		return user_id_err
	}
	var req types.AttachPlatformSpecificJobsRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}

	defer r.Body.Close()
	if errors := req.Validate(); len(errors) > 0 {
		return InvalidRequestData(errors)
	}
	new_upwork_job, err := s.svc.AttachPlatformSpecificjobs(req, user_id, nil, r.Context())
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, new_upwork_job)
}

func (s *Server) GetUpworkJob(w http.ResponseWriter, r *http.Request) error {
	user_id, user_id_err := s.authenticator.GetIdFromRequest(r)
	if user_id_err != nil {
		return user_id_err
	}
	upwork_job_id := chi.URLParam(r, "upwork_job_id")
	job, err := s.svc.GetUpworkJob(upwork_job_id, user_id, r.Context())
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, job)
}

func (s *Server) GetUpworkJobEmbeddingData(w http.ResponseWriter, r *http.Request) error {
	user_id, user_id_err := s.authenticator.GetIdFromRequest(r)
	if user_id_err != nil {
		return user_id_err
	}
	upwork_job_id := chi.URLParam(r, "upwork_job_id")
	job_id := chi.URLParam(r, "job_id")

	job, err := s.svc.GetUpworkJobEmbeddingData(upwork_job_id, user_id, r.Context())
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
	user_id, user_id_err := s.authenticator.GetIdFromRequest(r)
	if user_id_err != nil {
		return user_id_err
	}
	var req types.UpdateUpworkJobRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}

	defer r.Body.Close()
	if errors := req.Validate(); len(errors) > 0 {
		return InvalidRequestData(errors)
	}
	upwork_job, err := s.svc.UpdateUpworkJob(req, user_id, r.Context())
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, upwork_job)
}

func (s *Server) AddJobRankings(w http.ResponseWriter, r *http.Request) error {
	user_id, user_id_err := s.authenticator.GetIdFromRequest(r)
	if user_id_err != nil {
		return user_id_err
	}
	var req types.AddJobRankingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}

	defer r.Body.Close()
	if errors := req.Validate(); len(errors) > 0 {
		return InvalidRequestData(errors)
	}

	err := s.svc.AddJobRankings(req, user_id, r.Context())
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusCreated, nil)

}
