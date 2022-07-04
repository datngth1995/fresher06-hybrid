package common

import (
	"errors"
	"net/http"
)

func ErrorDBNotFound(err error) *ErrorResponse {
	return &ErrorResponse{ErrorRoot: err,
		ErrorUser: ErrorResponseToUser{
			StatusCode: http.StatusBadRequest,
			Message:    "Look up user failed",
		},
	}
}

func ErrorAuthentication(err error) *ErrorResponse {
	return &ErrorResponse{ErrorRoot: err,
		ErrorUser: ErrorResponseToUser{
			StatusCode: http.StatusBadRequest,
			Message:    "Authentication failed",
		},
	}
}

func ErrorTokenExpried(err error) *ErrorResponse {
	return &ErrorResponse{ErrorRoot: err,
		ErrorUser: ErrorResponseToUser{
			StatusCode: http.StatusBadRequest,
			Message:    "Token has been expired",
		},
	}
}

func ErrorTokenInvalid() *ErrorResponse {
	return &ErrorResponse{ErrorRoot: errors.New("Invalid token"),
		ErrorUser: ErrorResponseToUser{
			StatusCode: http.StatusBadRequest,
			Message:    "Token is invalid",
		},
	}
}

func ErrorUserDeleted() *ErrorResponse {
	return &ErrorResponse{ErrorRoot: errors.New("User deleted"),
		ErrorUser: ErrorResponseToUser{
			StatusCode: http.StatusBadRequest,
			Message:    "User not found",
		},
	}
}
