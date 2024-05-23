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
	s.Router.Route("/v1", func(v1_router chi.Router) {
		//public apis
		v1_router.Group(func(public_router chi.Router) {
			public_router.Route("/public", func(public_sub_router chi.Router) {
				public_sub_router.Route("/users", func(users_router chi.Router) {
					users_router.Post("/", Make(s.CreateUser))
					users_router.Post("/update", Make(s.UpdateUser))
				})
			})
		})
		//private apis
		v1_router.Group(func(private_router chi.Router) {
			private_router.Use(clerkhttp.RequireHeaderAuthorization())
			private_router.Route("/private", func(private_sub_router chi.Router) {
				private_sub_router.Route("/jobs", func(jobs_router chi.Router) {
					jobs_router.Post("/", Make(s.CreateJob))
					jobs_router.Route("/{job_id}", func(job_id_router chi.Router) {
						job_id_router.Post("/freelancers", Make(s.CreateFreelancers))
						job_id_router.Get("/", Make(s.GetJobByID))
					})
				})
			})
		})

	})
	return http.ListenAndServe(s.Port, s.Router)
}
