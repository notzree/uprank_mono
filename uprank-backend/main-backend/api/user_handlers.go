package api

import (
	"encoding/json"
	"log"
	"net/http"

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
	log.Printf("Creating user")
	new_user, err := s.ent.User.Create().
		SetID(req.User.ID).
		SetFirstName(req.User.FirstName).
		SetCompanyName(req.User.CompanyName).
		SetEmail(req.User.Email).
		Save(r.Context())
	log.Printf("Finished creating user")
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, new_user)

}
