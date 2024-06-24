package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	auth "github.com/notzree/uprank-backend/main-backend/authenticator"
	svc "github.com/notzree/uprank-backend/main-backend/service"
)

// todo: remove hard dependenceis on ent and clerk
type Server struct {
	Port          string
	authenticator auth.Authenticator
	Router        *chi.Mux
	svc           svc.Servicer
}

func NewServer(listen_addr string, router *chi.Mux, authenticator auth.Authenticator, servicer svc.Servicer) *Server {
	return &Server{
		Port:          listen_addr,
		Router:        router,
		authenticator: authenticator,
		svc:           servicer,
	}
}

func (s *Server) Start() error {
	s.Router.Route("/v1", func(v1_router chi.Router) {
		//public apis
		v1_router.Group(func(public_router chi.Router) {
			public_router.Route("/public", func(public_sub_router chi.Router) {
				public_sub_router.Route("/user", func(users_router chi.Router) {
					users_router.Post("/", Make(s.CreateUser))
					users_router.Post("/update", Make(s.UpdateUser))
				})
			})
		})
		v1_router.Post("/test", Make(s.TestRanking))
		v1_router.Get("/test/jobs/{job_id}/upwork/{upwork_job_id}", Make(s.GetUpworkJob))
		//private apis
		v1_router.Group(func(private_router chi.Router) {
			private_router.Use(func(next http.Handler) http.Handler {
				return s.authenticator.AuthenticationMiddleware(next)
			})
			//Job will have the main create job function and also probably end up housing apis to attach other platforms, crosspost, etc.
			//also todo: Refactor the apis to use the user edge instead of having to query the job then the user
			private_router.Route("/private", func(private_sub_router chi.Router) {
				private_sub_router.Route("/jobs", func(jobs_router chi.Router) {
					jobs_router.Post("/", Make(s.CreateJob))
					jobs_router.Route("/{job_id}", func(job_id_router chi.Router) {
						jobs_router.Post("/attach", Make(s.AttachUpworkJob)) //todo: move this attach upwork job into a general attach func
					})
					jobs_router.Route("/upwork", func(upwork_router chi.Router) {
						upwork_router.Route("/{upwork_job_id}", func(upwork_job_id_router chi.Router) {
							upwork_job_id_router.Get("/", Make(s.GetUpworkJob))
							upwork_job_id_router.Get("/all_data", Make(s.GetUpworkJobWithAllFreelancerData))
							upwork_job_id_router.Route("/freelancers", func(upwork_freelancers_router chi.Router) {
								upwork_freelancers_router.Post("/", Make(s.CreateUpworkFreelancers))       //create freelancers
								upwork_freelancers_router.Post("/update", Make(s.UpdateUpworkFreelancers)) //update freelancers
							})
						})
					})

				})
			})
		})
	})
	s.Router.Get("/healthz", Make(s.HealthCheck))
	return http.ListenAndServe(s.Port, s.Router)
}
