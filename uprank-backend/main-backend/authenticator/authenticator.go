package authenticator

import (
	"errors"
	"net/http"
	"strings"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
)

type Authenticator interface {
	AuthenticationMiddleware(next http.Handler) http.Handler
	GetIdFromRequest(r *http.Request) (string, error)
	GetToken(r *http.Request) (string, error)
}

type ClerkAuthenticator struct {
	clerk_secret_key     string
	inference_server_url string
}

func NewClerkAuthenticator(clerk_secret_key string, inference_server_url string) *ClerkAuthenticator {
	clerk.SetKey(clerk_secret_key)
	return &ClerkAuthenticator{clerk_secret_key: clerk_secret_key, inference_server_url: inference_server_url}
}

// func (c *ClerkAuthenticator) AuthenticationMiddleware(next http.Handler) http.Handler {
// 	return clerkhttp.WithHeaderAuthorization()(next)
// }

func (c *ClerkAuthenticator) AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Origin") == c.inference_server_url {
			//request made from inference server
			//TODO BEFORE PROD: Use api key or smth to ensure req is from inference server
			//maybe acc use jwt
			next.ServeHTTP(w, r)
		} else {
			token, err := c.GetToken(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			claims, err := jwt.Verify(r.Context(), &jwt.VerifyParams{
				Token: token,
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			newCtx := clerk.ContextWithSessionClaims(r.Context(), claims)
			next.ServeHTTP(w, r.WithContext(newCtx))
		}
	})

}

func (c *ClerkAuthenticator) GetIdFromRequest(r *http.Request) (string, error) {
	ctx := r.Context()
	claims, exist := clerk.SessionClaimsFromContext(ctx)
	if !exist {
		// try to grab id from request:
		id := r.Header.Get("User_id")
		if id == "" {
			return "", errors.New("user_id is required")
		}
		return id, nil
	}
	return claims.Subject, nil
}

func (c *ClerkAuthenticator) GetToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is required")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == authHeader {
		return "", errors.New("invalid authorization header format")
	}

	return token, nil
}
