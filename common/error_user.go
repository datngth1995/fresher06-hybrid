package common

import "net/http"

func ErrorDBNotFound(err error) *ErrorResponse {
	return &ErrorResponse{RootError: err,
		StatusCode: http.StatusBadRequest,
		Message:    "Look up user failed"}
}

func ErrorAuthentication(err error) *ErrorResponse {
	return &ErrorResponse{RootError: err,
		StatusCode: http.StatusBadRequest,
		Message:    "Authentication failed"}
}
