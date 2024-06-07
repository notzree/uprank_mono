package api

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/go-chi/chi/v5"
	auth "github.com/notzree/uprank-backend/main-backend/authenticator"
	"github.com/notzree/uprank-backend/main-backend/ent"
)

// todo: remove hard dependenceis on ent and clerk
type Server struct {
	Port                 string
	authenticator        auth.Authenticator
	scraper_queue_url    string
	Router               *chi.Mux
	ent                  *ent.Client
	scraper_queue_client *sqs.Client
}

func NewServer(listen_addr string, router *chi.Mux, ent *ent.Client, scraper_queue_url string, scraper_queue_client *sqs.Client, authenticator auth.Authenticator) *Server {
	return &Server{
		Port:                 listen_addr,
		Router:               router,
		ent:                  ent,
		scraper_queue_url:    scraper_queue_url,
		scraper_queue_client: scraper_queue_client,
		authenticator:        authenticator,
	}
}

func (s *Server) Start() error {
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
			private_router.Use(func(next http.Handler) http.Handler {
				return s.authenticator.AuthenticationMiddleware(next)
			})
			private_router.Route("/private", func(private_sub_router chi.Router) {
				private_sub_router.Route("/jobs", func(jobs_router chi.Router) {
					jobs_router.Post("/", Make(s.CreateJob))
					jobs_router.Route("/{job_id}", func(job_id_router chi.Router) {
						job_id_router.Post("/freelancers", Make(s.CreateFreelancers))
						job_id_router.Post("/freelancers/update", Make(s.UpdateFreelancers))
						job_id_router.Get("/", Make(s.GetJobByID))
					})
				})
			})
		})
	})
	s.Router.Get("/healthz", Make(s.HealthCheck))
	return http.ListenAndServe(s.Port, s.Router)
}
