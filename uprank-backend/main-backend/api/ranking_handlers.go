package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (s *Server) TestRanking(w http.ResponseWriter, r *http.Request) error {
	var req TestRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}
	err := s.svc.SendRankingrequest(req.JobID, req.UserID, r.Context())
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, nil)

}

type TestRequest struct {
	JobID  uuid.UUID `json:"job_id"`
	UserID string    `json:"user_id"`
}
