package authenticator

import (
	"context"
	"errors"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
)

type Authenticator interface {
	AuthenticationMiddleware(next http.Handler) http.Handler

	GetIdFromContext(ctx context.Context) (string, error)
}

type ClerkAuthenticator struct {
	clerk_secret_key string
}

func NewClerkAuthenticator(clerk_secret_key string) *ClerkAuthenticator {
	clerk.SetKey(clerk_secret_key)
	return &ClerkAuthenticator{clerk_secret_key: clerk_secret_key}
}

func (c *ClerkAuthenticator) AuthenticationMiddleware(next http.Handler) http.Handler {
	return clerkhttp.RequireHeaderAuthorization()(next)
}
func (c *ClerkAuthenticator) GetIdFromContext(ctx context.Context) (string, error) {
	claims, exist := clerk.SessionClaimsFromContext(ctx)
	if !exist {
		return "", errors.New("no claims found in context")
	}
	return claims.Subject, nil
}
