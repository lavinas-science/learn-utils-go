package rest_errors

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRestErr(t *testing.T) {
	message := "Message"
	status := http.StatusBadRequest
	error := "dada"
	e := NewRestErr(message, status, error)
	assert.NotNil(t, e)
	assert.EqualValues(t, e.Message(), message)
	assert.EqualValues(t, e.Error(), fmt.Sprintf("message: %s - status: %d - error: %s", message, status, error))
	assert.EqualValues(t, e.Status(), status)
}

func TestNewBadRequestError(t *testing.T) {
	message := "Message"
	status := http.StatusBadRequest
	error := "bad_request"
	e := NewBadRequestError(message)
	assert.NotNil(t, e)
	assert.EqualValues(t, e.Message(), message)
	assert.EqualValues(t, e.Error(), fmt.Sprintf("message: %s - status: %d - error: %s", message, status, error))
	assert.EqualValues(t, e.Status(), status)
}

func TestNewNotFoundError(t *testing.T) {
	message := "Message"
	status := http.StatusNotFound
	error := "not_found"
	e := NewNotFoundError(message)
	assert.NotNil(t, e)
	assert.EqualValues(t, e.Message(), message)
	assert.EqualValues(t, e.Error(), fmt.Sprintf("message: %s - status: %d - error: %s", message, status, error))
	assert.EqualValues(t, e.Status(), status)
}

func TestNewInternalServerError(t *testing.T) {
	message := "Message"
	status := http.StatusInternalServerError
	error := "internal_server_error"
	e := NewInternalServerError(message)
	assert.NotNil(t, e)
	assert.EqualValues(t, e.Message(), message)
	assert.EqualValues(t, e.Error(), fmt.Sprintf("message: %s - status: %d - error: %s", message, status, error))
	assert.EqualValues(t, e.Status(), status)
}

func TestNewNotImplementedError(t *testing.T) {
	message := "Message"
	status := http.StatusNotImplemented
	error := "not_implemented"
	e := NewNotImplementedError(message)
	assert.NotNil(t, e)
	assert.EqualValues(t, e.Message(), message)
	assert.EqualValues(t, e.Error(), fmt.Sprintf("message: %s - status: %d - error: %s", message, status, error))
	assert.EqualValues(t, e.Status(), status)
}

func TestNewUnauthorizedError(t *testing.T) {
	message := "Message"
	status := http.StatusUnauthorized
	error := "unauthorized"
	e := NewUnauthorizedError(message)
	assert.NotNil(t, e)
	assert.EqualValues(t, e.Message(), message)
	assert.EqualValues(t, e.Error(), fmt.Sprintf("message: %s - status: %d - error: %s", message, status, error))
	assert.EqualValues(t, e.Status(), status)
}
