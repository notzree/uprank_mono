package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/notzree/uprank-backend/main-backend/types"
)

func (s *Server) TestRanking(w http.ResponseWriter, r *http.Request) error {

	var req TestRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}
	data := types.RankJobRequest{
		Job_id:            req.JobID,
		User_id:           req.UserID,
		Short_lived_token: "poo",
		Platform:          req.Platform,
		Platform_id:       req.PlatformID,
	}
	err := s.svc.SendRankingrequest(data, r.Context())
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, nil)
}

type TestRequest struct {
	JobID      uuid.UUID `json:"job_id"`
	UserID     string    `json:"user_id"`
	Platform   string    `json:"platform"`
	PlatformID string    `json:"platform_id"`
}
