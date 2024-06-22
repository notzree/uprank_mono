package api

import (
	"encoding/json"
	"net/http"

	"github.com/notzree/uprank-backend/main-backend/types"
)

func (s *Server) CreateJob(w http.ResponseWriter, r *http.Request) error {
	user_id, user_id_err := s.authenticator.GetIdFromRequest(r)
	if user_id_err != nil {
		return user_id_err
	}
	var req types.CreateJobRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}

	if errors := req.Validate(); len(errors) > 0 {
		return InvalidRequestData(errors)
	}
	new_job, err := s.svc.CreateJob(req, user_id, r.Context())

	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, new_job)

}
