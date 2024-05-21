package api

import (
	"encoding/json"
	"net/http"

	"github.com/notzree/uprank-backend/main-backend/types"
)

func (s *Server) CreateJob(w http.ResponseWriter, r *http.Request) error {
	var req types.CreateJobRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}

	defer r.Body.Close()
	if errors := req.Validate(); len(errors) > 0 {
		return InvalidRequestData(errors)
	}
	new_user, err := s.ent.Job.Create().
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
		Save(r.Context())
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusCreated, new_user)
}
