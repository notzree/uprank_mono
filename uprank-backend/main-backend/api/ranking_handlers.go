package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/types"
)

func (s *Server) TestRanking(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	var req TestRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}
	data := types.RankJobRequest{
		Job_id:            req.JobID,
		User_id:           req.UserID,
		Short_lived_token: "unimplemented",
		Platform:          req.Platform,
		Platform_id:       req.PlatformID,
	}
	err := s.ranking.SendRankingrequest(ctx, data)
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

func (s *Server) AddJobRankings(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	_, user_id_err := s.authenticator.GetIdFromRequest(r)
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

	err := s.ranking.AddJobRankings(ctx, types.AddJobRankingRequest{
		Freelancer_ranking_data: req.Freelancer_ranking_data,
	})
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusCreated, nil)

}
