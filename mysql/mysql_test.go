package mysql

import (
	"errors"
	"testing"

	"github.com/lavinas-science/learn-utils-go/rest_errors"
	"github.com/stretchr/testify/assert"
)

func TestParseErrorParseError(t *testing.T) {
	err := errors.New("New error")
	p := ParseError(err)
	np := rest_errors.NewInternalServerError("Database error conversion - contact admin")
	assert.NotNil(t, p)
	assert.EqualValues(t, np, p)
}

func TestParseErrorNoRecord(t *testing.T) {
	err := errors.New("sql: no rows in result set")
	np := rest_errors.NewNotFoundError("No record found")
	p := ParseError(err)
	assert.NotNil(t, p)
	assert.EqualValues(t, np, p)
}
