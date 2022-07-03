package common

type ErrorResponse struct {
	StatusCode int    `json:"StatusCode"`
	Message    string `json:"message"`
	RootError  error  `json:"-"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}
