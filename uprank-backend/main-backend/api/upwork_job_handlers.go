package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/notzree/uprank-backend/main-backend/types"
)

func (s *Server) AttachUpworkJob(w http.ResponseWriter, r *http.Request) error {
	user_id, user_id_err := s.authenticator.GetIdFromContext(r.Context())
	if user_id_err != nil {
		return user_id_err
	}
	var req types.AttachUpworkJobRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}

	defer r.Body.Close()
	if errors := req.Validate(); len(errors) > 0 {
		return InvalidRequestData(errors)
	}
	new_upwork_job, err := s.svc.AttachUpworkJob(req, user_id, nil, r.Context())
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, new_upwork_job)
}

func (s *Server) GetUpworkJob(w http.ResponseWriter, r *http.Request) error {
	user_id, user_id_err := s.authenticator.GetIdFromContext(r.Context())
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
