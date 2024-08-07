package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type APIError struct {
	StatusCode int `json:"statusCode"`
	Msg        any `json:"msg"`
}

func (e APIError) Error() string {
	if msg, ok := e.Msg.(string); ok && msg != "" {
		return fmt.Sprintf("api error: %d: %s", e.StatusCode, msg)
	}
	return fmt.Sprintf("api error: %d", e.StatusCode)
}

func NewAPIError(statusCode int, err error) APIError {
	return APIError{StatusCode: statusCode, Msg: err.Error()}
}

func InvalidRequestData(errors map[string]interface{}) APIError {
	return APIError{
		StatusCode: http.StatusUnprocessableEntity,
		Msg:        errors,
	}
}

func InvalidJSON() APIError {
	return NewAPIError(http.StatusBadRequest, fmt.Errorf("invalid JSON request data"))
}

func ResourceMisMatch() APIError {
	return NewAPIError(http.StatusBadRequest, fmt.Errorf("mismatch between user_id and resource"))
}

func NotFound() APIError {
	return NewAPIError(http.StatusNotFound, fmt.Errorf("resource not found"))
}

type APIFunc func(w http.ResponseWriter, r *http.Request) error

// converts an APIFunc (a function that returns an error) into a function that does not return an error http.HandlerFunc
// Will either respond with an API error or an internal server error and log the error
// If the function does not return an APIError, it will be treated as an internal server error (not sent to user)
func Make(handler APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			if apiErr, ok := err.(APIError); ok {
				writeJSON(w, apiErr.StatusCode, apiErr.Msg)
			} else {
				errResp := map[string]any{
					"statusCode": http.StatusInternalServerError,
					"msg":        "internal server error",
				}
				writeJSON(w, http.StatusInternalServerError, errResp)
			}
			slog.Error("HTTP API error", "err", err.Error(), "path", r.URL.Path)
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
