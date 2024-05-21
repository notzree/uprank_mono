package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/notzree/uprank-backend/main-backend/types"
)

// Creates a user in the database after they have signed up w/ clerk + completed onboarding
func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) error {
	var req types.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}

	defer r.Body.Close()
	if errors := req.Validate(); len(errors) > 0 {
		return InvalidRequestData(errors)
	}
	new_user, err := s.ent.User.Create().
		SetID(req.User.ID).
		SetFirstName(req.User.FirstName).
		SetCompanyName(req.User.CompanyName).
		SetEmail(req.User.Email).
		Save(r.Context())
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusCreated, new_user)
}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	//Clerk sends the request via webhooks so the body is guaranteed to be well formatted
	var req types.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}
	defer r.Body.Close()

	updated_user, err := s.ent.User.
		UpdateOneID(req.ClerkUserData.ID).
		SetFirstName(req.ClerkUserData.FirstName).
		SetLastLogin(time.Unix(req.ClerkUserData.LastSignInAt, 0)). //Convert timestamp into time.Time
		SetUpdatedAt(time.Unix(req.ClerkUserData.UpdatedAt, 0)).
		Save(r.Context())

	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, updated_user)
}
