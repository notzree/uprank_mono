package api

import "net/http"

func (s *Server) HealthCheck(w http.ResponseWriter, r *http.Request) error {

	if clerkErr := checkClerk(); clerkErr != nil {
		return clerkErr
	}
	_, err := s.ent.User.Query().First(r.Context())
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, "healthy")
}

func checkClerk() error {
	_, err := http.Get("https://upright-crow-79.accounts.dev")
	if err != nil {
		return err
	}
	return nil

}