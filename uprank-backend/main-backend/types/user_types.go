package types

import (
	"strings"

	"github.com/notzree/uprank-backend/main-backend/ent"
)

type CreateUserRequest struct {
	User                 *ent.User `json:"user"`
	Completed_onboarding bool      `json:"completed_onboarding"`
}

func (req *CreateUserRequest) Validate() map[string]string {
	errors := make(map[string]string)

	if req.User == nil {
		errors["user"] = "User cannot be nil"
		return errors
	}

	// Check required fields in the User struct
	if strings.TrimSpace(req.User.FirstName) == "" {
		errors["first_name"] = "First name cannot be empty"
	}
	if strings.TrimSpace(req.User.CompanyName) == "" {
		errors["company_name"] = "Company name cannot be empty"
	}
	if strings.TrimSpace(req.User.Email) == "" {
		errors["email"] = "Email cannot be empty"
	}
	if !strings.Contains(req.User.Email, "@") {
		errors["email"] = "Email must contain @"
	}

	return errors
}
