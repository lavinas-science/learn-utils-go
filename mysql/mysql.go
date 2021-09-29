package mysql

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/lavinas-science/learn-utils-go/rest_errors"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) rest_errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return rest_errors.NewNotFoundError("No record found")
		}
		return rest_errors.NewInternalServerError("Database error conversion - contact admin")
	}
	switch sqlErr.Number {
	case 1062:
		return rest_errors.NewBadRequestError("Duplicated data")
	}
	return rest_errors.NewInternalServerError("Database error - contact admin")
}
