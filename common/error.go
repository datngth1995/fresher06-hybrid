package common

import "net/http"

type ErrorResponseToUser struct {
	StatusCode int    `json:"StatusCode"`
	Message    string `json:"message"`
}

type ErrorResponse struct {
	ErrorRoot error `json:"-"`
	ErrorUser ErrorResponseToUser
}

func (e *ErrorResponse) Error() string {
	return e.ErrorUser.Message
}

func ErrorInternal(err error) *ErrorResponse {
	return &ErrorResponse{ErrorRoot: err,
		ErrorUser: ErrorResponseToUser{
			StatusCode: http.StatusBadRequest,
			Message:    "Internal Error",
		},
	}
}
