package server

import (
	"fmt"
	"log/slog"
)

type QError struct {
	Msg any `json:"msg"`
}

func (e QError) Error() string {
	return fmt.Sprintf("api error: %s:", e.Msg)
}

func NewQError(err error) QError {
	return QError{Msg: err.Error()}
}

type ServiceError struct {
	Msg          any    `json:"msg"`
	FunctionName string `json:"function_name"`
}

func (e ServiceError) Error() string {
	return fmt.Sprintf(" %s | %s:", e.FunctionName, e.Msg)
}

func NewServiceError(err error, function_name string) ServiceError {
	return ServiceError{
		Msg:          err.Error(),
		FunctionName: function_name,
	}
}

func HandleError(err error) {
	if q_err, ok := err.(QError); ok {
		slog.Error("Queue Error", "msg", q_err.Msg)
	} else if s_err, ok := err.(ServiceError); ok {
		slog.Error("Service Error", "msg", s_err.Msg)
	} else {
		slog.Error("internal server error", "err", err.Error())

	}
	slog.Error("function error", "err", err.Error())
}
