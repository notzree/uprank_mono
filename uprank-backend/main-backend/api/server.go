package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/notzree/uprank-backend/main-backend/ent"
)

type Server struct {
	Port   string
	Router *chi.Mux
	ent    *ent.Client
	//remember to add clerk client here or in the middleware
}

func NewServer(listen_addr string, router *chi.Mux, Db *ent.Client) *Server {
	return &Server{
		Port:   listen_addr,
		Router: router,
		ent:    Db,
	}
}

func (s *Server) Start() error {
	//public apis
	s.Router.Group(func(r chi.Router) {
		r.Route("/public", func(r chi.Router) {
			r.Post("/users", Make(s.CreateUser))
		})

	})
	//private apis
	s.Router.Group(func(r chi.Router) {
		// r.Use(AuthMiddleware)  todo: implement AuthMiddleware
	})
	return http.ListenAndServe(s.Port, s.Router)
}
