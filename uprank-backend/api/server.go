package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	listen_addr string
	router *chi.Mux
}

func NewServer (listen_addr string, router *chi.Mux) *Server {
	return &Server{
		listen_addr: listen_addr,
		router: router,
	}
}

func (s *Server) Start() error {

	
	return http.ListenAndServe(s.listen_addr, s.router)
}

