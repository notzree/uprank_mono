package service

import (
	"context"
	"time"

	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/types"
)

type UserService struct {
	ent *ent.Client
}

type NewUserServiceParams struct {
	Ent *ent.Client
}

func NewUserService(params NewUserServiceParams) UserService {
	return UserService{
		ent: params.Ent,
	}
}

func (u *UserService) CreateUser(ctx context.Context, data types.CreateUserRequest) (*ent.User, error) {
	new_user, err := u.ent.User.Create().
		SetID(data.User.ID).
		SetFirstName(data.User.FirstName).
		SetCompanyName(data.User.CompanyName).
		SetEmail(data.User.Email).
		Save(ctx)
	return new_user, err
}

func (u *UserService) UpdateUser(ctx context.Context, data types.UpdateUserRequest) (*ent.User, error) {
	updated_user, err := u.ent.User.
		UpdateOneID(data.ClerkUserData.ID).
		SetFirstName(data.ClerkUserData.FirstName).
		SetLastLogin(time.Unix(data.ClerkUserData.LastSignInAt, 0)). //Convert timestamp into time.Time
		SetUpdatedAt(time.Unix(data.ClerkUserData.UpdatedAt, 0)).
		Save(ctx)
	return updated_user, err
}
