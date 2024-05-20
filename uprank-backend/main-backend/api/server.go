package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	ListenAddr string
	Router     *chi.Mux
}

func NewServer(listen_addr string, router *chi.Mux) *Server {
	return &Server{
		ListenAddr: listen_addr,
		Router:     router,
	}
}

func (s *Server) Start() error {

	return http.ListenAndServe(s.ListenAddr, s.Router)
}
