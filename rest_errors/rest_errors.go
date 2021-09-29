package rest_errors

import (
	"fmt"
	"net/http"
)

type RestErr interface {
	Error() string
	Message() string
	Status() int
}

type restErr struct {
	message string `json:"message"`
	status  int    `json:"status"`
	error   string `json:"error"`
}

func (e restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error %s ", e.message, e.status, e.error)
}

func (e restErr) Message() string {
	return e.message
}

func (e restErr) Status() int {
	return e.status
}

func NewRestErr(message string, status int, error string) RestErr {
	return restErr{message, status, error}
}

func NewBadRequestError(message string) RestErr {
	return NewRestErr(message, http.StatusBadRequest, "bad_request")
}

func NewNotFoundError(message string) RestErr {
	return NewRestErr(message, http.StatusNotFound, "not_found")
}

func NewInternalServerError(message string) RestErr {
	return NewRestErr(message, http.StatusInternalServerError, "internal_server_error")
}

func NewNotImplementedError(message string) RestErr {
	return NewRestErr(message, http.StatusNotImplemented, "not_implemented")
}

func NewUnauthorizedError(message string) RestErr {
	return NewRestErr(message, http.StatusUnauthorized, "unauthorized")
}
