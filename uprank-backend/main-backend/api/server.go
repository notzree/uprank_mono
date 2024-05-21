package api

import (
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/go-chi/chi/v5"
	"github.com/notzree/uprank-backend/main-backend/ent"
)

type Server struct {
	Port             string
	Router           *chi.Mux
	ent              *ent.Client
	Clerk_secret_key string
}

func NewServer(listen_addr string, router *chi.Mux, ent *ent.Client, clerk_secret_key string) *Server {
	return &Server{
		Port:             listen_addr,
		Router:           router,
		ent:              ent,
		Clerk_secret_key: clerk_secret_key,
	}
}

func (s *Server) Start() error {
	clerk.SetKey(s.Clerk_secret_key)
	s.Router.Route("/v1", func(r chi.Router) {
		//public apis
		r.Group(func(r chi.Router) {
			r.Route("/public", func(r chi.Router) {
				r.Post("/users", Make(s.CreateUser))
				r.Post("/users/update", Make(s.UpdateUser))
			})
		})
		//private apis
		r.Group(func(r chi.Router) {
			r.Use(clerkhttp.RequireHeaderAuthorization())
			r.Route("/private", func(r chi.Router) {

			})
		})

	})
	return http.ListenAndServe(s.Port, s.Router)
}
