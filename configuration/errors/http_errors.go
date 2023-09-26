package http_errors

import "net/http"

type HttpErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *HttpErr) Error() string {
	return r.Message
}

func NewBadRequestError(message string) *HttpErr {
	return &HttpErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewUnauthorizedRequestError(message string) *HttpErr {
	return &HttpErr{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *HttpErr {
	return &HttpErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *HttpErr {
	return &HttpErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *HttpErr {
	return &HttpErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}
