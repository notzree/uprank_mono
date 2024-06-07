package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/notzree/uprank-backend/main-backend/ent/job"
	"github.com/notzree/uprank-backend/main-backend/ent/user"
	"github.com/notzree/uprank-backend/main-backend/types"
)

func (s *Server) CreateJob(w http.ResponseWriter, r *http.Request) error {
	user_id, user_id_err := s.authenticator.GetIdFromContext(r.Context())
	if user_id_err != nil {
		return user_id_err
	}
	var req types.CreateJobRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}

	defer r.Body.Close()
	if errors := req.Validate(); len(errors) > 0 {
		return InvalidRequestData(errors)
	}
	new_job, err := s.ent.Job.Create().
		SetID(req.Id).
		SetTitle(req.Title).
		SetLocation(req.Location).
		SetDescription(req.Description).
		SetSkills(req.Skills).
		SetExperienceLevel(req.Experience_level).
		SetHourly(req.Hourly).
		SetFixed(req.Fixed).
		SetHourlyRate(req.Hourly_rate).
		SetFixedRate(req.Fixed_rate).
		SetUserID(user_id).
		Save(r.Context())
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusCreated, new_job)
}

func (s *Server) GetJobByID(w http.ResponseWriter, r *http.Request) error {
	user_id, user_id_err := s.authenticator.GetIdFromContext(r.Context())
	if user_id_err != nil {
		return user_id_err
	}
	job_id := chi.URLParam(r, "job_id")
	job, err := s.ent.Job.Query().
		Where(
			job.IDEQ(job_id),
			job.HasUserWith(user.IDEQ(user_id)),
		).
		WithFreelancers().
		Only(context.Background())

	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, job)
}
