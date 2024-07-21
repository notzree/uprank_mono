package api

import (
	"encoding/json"
	"net/http"

	"github.com/notzree/uprank_mono/uprank-backend/main-backend/types"
)

// Creates a user in the database after they have signed up w/ clerk + completed onboarding
func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	var req types.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}
	defer r.Body.Close()
	if errors := req.Validate(); len(errors) > 0 {
		return InvalidRequestData(errors)
	}

	new_user, err := s.user.CreateUser(ctx, req)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusCreated, new_user)
}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	//Clerk sends the request via webhooks so the body is guaranteed to be well formatted
	var req types.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}
	defer r.Body.Close()

	updated_user, err := s.user.UpdateUser(ctx, req)

	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, updated_user)
}
