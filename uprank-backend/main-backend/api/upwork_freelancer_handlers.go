package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/notzree/uprank-backend/main-backend/types"
)

func (s *Server) CreateUpworkFreelancers(w http.ResponseWriter, r *http.Request) error {
	user_id, user_id_err := s.authenticator.GetIdFromContext(r.Context())
	if user_id_err != nil {
		return user_id_err
	}
	upwork_job_id := chi.URLParam(r, "upwork_job_id")
	var req []types.CreateUpworkFreelancerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("failed to decode request", "err", err)
		return InvalidJSON()
	}
	defer r.Body.Close()

	for _, freelancer := range req {
		if errors := freelancer.Validate(); len(errors) > 0 {
			return InvalidRequestData(errors)
		}
	}

	freelancers, err := s.svc.CreateUpworkFreelancer(req, user_id, upwork_job_id, r.Context())
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, freelancers)
}

func (s *Server) UpdateUpworkFreelancers(w http.ResponseWriter, r *http.Request) error {
	user_id, user_id_err := s.authenticator.GetIdFromContext(r.Context())
	if user_id_err != nil {
		return user_id_err
	}
	upwork_job_id := chi.URLParam(r, "upwork_job_id")

	var (
		req []types.CreateUpworkFreelancerRequest
	)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("failed to decode request", "err", err)
		return InvalidJSON()
	}
	defer r.Body.Close()

	for _, freelancer := range req {
		if errors := freelancer.Validate(); len(errors) > 0 {
			return InvalidRequestData(errors)
		}
	}

	created_freelancer_count, updated_freelancer_count, deleted_freelancer_count, err := s.svc.UpdateUpworkFreelancer(req, user_id, upwork_job_id, r.Context())
	if err != nil {
		return err
	}

	update_freelancer_response := types.UpdateFreelancerResponse{
		CreatedFreelancerCount: created_freelancer_count,
		UpdatedFreelancerCount: updated_freelancer_count,
		DeletedFreelancerCount: deleted_freelancer_count,
	}

	return writeJSON(w, http.StatusCreated, update_freelancer_response)
}
