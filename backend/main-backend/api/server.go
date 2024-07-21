package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	auth "github.com/notzree/uprank_mono/uprank-backend/main-backend/authenticator"
	svc "github.com/notzree/uprank_mono/uprank-backend/main-backend/service"
)

// todo: remove hard dependenceis on ent and clerk
type Server struct {
	Port              string
	authenticator     auth.Authenticator
	router            *chi.Mux
	user              svc.UserService
	job               svc.JobService
	upwork_job        svc.UpworkJobService
	upwork_freelancer svc.UpworkFreelancerService
	ranking           svc.RankingService
}

type NewServerParams struct {
	Authenticator     auth.Authenticator
	Router            *chi.Mux
	User              svc.UserService
	Job               svc.JobService
	Upwork_job        svc.UpworkJobService
	Upwork_freelancer svc.UpworkFreelancerService
	Ranking           svc.RankingService
}

type Option func(s *Server)

func WithPort(port string) Option {
	return func(s *Server) {
		s.Port = port
	}
}

func NewServer(params NewServerParams, opts ...Option) (*Server, error) {

	s := &Server{
		router:        params.Router,
		authenticator: params.Authenticator,
		user:          params.User,
		job:           params.Job,
		upwork_job:    params.Upwork_job,
		ranking:       params.Ranking,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s, nil
}

func (s *Server) Start() error {
	s.router.Route("/v1", func(v1_router chi.Router) {
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
					jobs_router.Get("/", Make(s.GetJobs))
					jobs_router.Route("/{job_id}", func(job_id_router chi.Router) {
						job_id_router.Post("/attach", Make(s.AttachPlatformSpecificJobs)) //todo: move this attach upwork job into a general attach func
						job_id_router.Get("/", Make(s.GetJobById))
						job_id_router.Route("/upwork", func(upwork_router chi.Router) {
							upwork_router.Route("/{upwork_job_id}", func(upwork_job_id_router chi.Router) {
								upwork_job_id_router.Get("/", Make(s.GetUpworkJob))
								upwork_job_id_router.Post("/rank", Make(s.AddJobRankings))
								upwork_job_id_router.Post("/update", Make(s.UpdateUpworkJob)) //todo: make this api take in an array instead so we can batch update
								upwork_job_id_router.Route("/freelancers", func(upwork_freelancers_router chi.Router) {
									upwork_freelancers_router.Post("/", Make(s.UpsertUpworkFreelancer))       //create freelancers
									upwork_freelancers_router.Post("/upsert", Make(s.UpsertUpworkFreelancer)) //TODO: rename the route to /upsert
									upwork_freelancers_router.Post("/update", Make(s.UpdateUpworkFreelancer))
								})
								upwork_job_id_router.Route("/embeddings", func(upwork_job_id_embedding_router chi.Router) {
									upwork_job_id_embedding_router.Get("/job_data", Make(s.GetUpworkJobEmbeddingData))
								})
							})
						})
					})

				})
			})
		})
	})
	s.router.Get("/healthz", Make(s.HealthCheck))
	return http.ListenAndServe(s.Port, s.router)
}