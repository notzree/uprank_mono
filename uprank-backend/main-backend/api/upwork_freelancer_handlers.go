package api

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/types"
)

func (s *Server) CreateUpworkFreelancers(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	userId, userIdErr := s.authenticator.GetIdFromRequest(r)
	if userIdErr != nil {
		return userIdErr
	}
	upworkJobId := chi.URLParam(r, "upwork_job_id")
	var req []types.CreateUpworkFreelancerData

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

	freelancers, err := s.upwork_freelancer.CreateUpworkFreelancer(ctx, types.CreateUpworkFreelancerRequest{
		UserId:      userId,
		UpworkJobId: upworkJobId,
		Data:        req,
	})
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, freelancers)
}

func (s *Server) UpsertUpworkFreelancer(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	userId, user_id_err := s.authenticator.GetIdFromRequest(r)
	if user_id_err != nil {
		return user_id_err
	}
	upworkJobId := chi.URLParam(r, "upwork_job_id")

	var (
		req []types.CreateUpworkFreelancerData
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

	response, err := s.upwork_freelancer.UpsertUpworkFreelancer(ctx, types.CreateUpworkFreelancerRequest{
		UserId:      userId,
		UpworkJobId: upworkJobId,
		Data:        req,
	})
	if err != nil {
		return err
	}

	update_freelancer_response := types.UpsertFreelancerResponse{
		CreatedFreelancerCount: response.CreatedFreelancerCount,
		UpdatedFreelancerCount: response.UpdatedFreelancerCount,
		DeletedFreelancerCount: response.DeletedFreelancerCount,
	}

	return writeJSON(w, http.StatusCreated, update_freelancer_response)
}

func (s *Server) UpdateUpworkFreelancer(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	//TODO: FIGURE OUT WHY THIS SHIT TAKES 14 SECONDS
	//oh i know why it takes 14 seconds CUZ ENT IS TRASH AT THIS
	//CHANGE THIS TO RAW SQL
	userId, user_id_err := s.authenticator.GetIdFromRequest(r)
	if user_id_err != nil {
		return user_id_err
	}
	upworkJobId := chi.URLParam(r, "upwork_job_id")

	var (
		req []types.UpdateUpworkFreelancerData
	)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("failed to decode request", "err", err)
		return InvalidJSON()
	}
	defer r.Body.Close()

	for _, freelancer := range req {
		if errors := freelancer.Validate(); len(errors) > 0 {
			log.Println(errors)
			return InvalidRequestData(errors)
		}
	}
	start := time.Now()
	updated_ids, err := s.upwork_freelancer.UpdateUpworkFreelancer(ctx, types.UpdateUpworkFreelancerRequest{
		UserId:      userId,
		UpworkJobId: upworkJobId,
		Data:        req,
	})
	elapsed := time.Since(start)
	log.Default().Println("this shit took ", elapsed)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, updated_ids)

}
