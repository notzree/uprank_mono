package types

import (
	"strings"
	"time"

	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent"
)

type CreateUserRequest struct {
	User                 *ent.User `json:"user"`
	Completed_onboarding bool      `json:"completed_onboarding"`
}

func (req *CreateUserRequest) Validate() map[string]interface{} {
	errors := make(map[string]interface{})

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

type UpdateUserRequest struct {
	ClerkUserData ClerkUserData `json:"data"`
	Object        string        `json:"object"`
	Type          string        `json:"type"`
}
type ClerkUserData struct {
	Birthday              string                 `json:"birthday"`
	CreatedAt             int64                  `json:"created_at"`
	EmailAddresses        []EmailAddress         `json:"email_addresses"`
	ExternalAccounts      []interface{}          `json:"external_accounts"`
	ExternalID            interface{}            `json:"external_id"`
	FirstName             string                 `json:"first_name"`
	Gender                string                 `json:"gender"`
	ID                    string                 `json:"id"`
	ImageURL              string                 `json:"image_url"`
	LastName              *string                `json:"last_name"`
	LastSignInAt          int64                  `json:"last_sign_in_at"`
	Object                string                 `json:"object"`
	PasswordEnabled       bool                   `json:"password_enabled"`
	PhoneNumbers          []interface{}          `json:"phone_numbers"`
	PrimaryEmailAddressID string                 `json:"primary_email_address_id"`
	PrimaryPhoneNumberID  *string                `json:"primary_phone_number_id"`
	PrimaryWeb3WalletID   *string                `json:"primary_web3_wallet_id"`
	PrivateMetadata       map[string]interface{} `json:"private_metadata"`
	ProfileImageURL       string                 `json:"profile_image_url"`
	PublicMetadata        map[string]interface{} `json:"public_metadata"`
	TwoFactorEnabled      bool                   `json:"two_factor_enabled"`
	UnsafeMetadata        map[string]interface{} `json:"unsafe_metadata"`
	UpdatedAt             int64                  `json:"updated_at"`
	Username              *string                `json:"username"`
	Web3Wallets           []interface{}          `json:"web3_wallets"`
}

// EmailAddress represents the email address structure
type EmailAddress struct {
	EmailAddress string        `json:"email_address"`
	ID           string        `json:"id"`
	LinkedTo     []interface{} `json:"linked_to"`
	Object       string        `json:"object"`
	Reserved     bool          `json:"reserved"`
	Verification Verification  `json:"verification"`
}

// Verification represents the verification structure
type Verification struct {
	Attempts *int       `json:"attempts"`
	ExpireAt *time.Time `json:"expire_at"`
	Status   string     `json:"status"`
	Strategy string     `json:"strategy"`
}
