package apis

import (
	"database/sql"
	"fmt"
	"net/http"
	"reflect"

	"github.com/duyanhitbe/library-golang/validations"
)

const (
	BadRequestMessage     = "Bad Request"
	InternalServerMessage = "Internal Server"
	ForbiddenMessage      = "Forbidden"
)

type ExceptionResponse struct {
	StatusCode int                           `json:"status_code"`
	Success    bool                          `json:"success"`
	Message    string                        `json:"message"`
	Error      string                        `json:"error"`
	Errors     []validations.ValidationError `json:"errors"`
}

func (server *HttpServer) ThrowException(exception *ExceptionResponse) {
	server.ctx.JSON(exception.StatusCode, exception)
}

func (server *HttpServer) ThrowBadRequestException(err error) {
	if validationErr, isValidationErr := validations.GetValidationError(err); isValidationErr {
		server.ThrowException(&ExceptionResponse{
			StatusCode: http.StatusBadRequest,
			Success:    false,
			Message:    BadRequestMessage,
			Error:      "Validation Error",
			Errors:     validationErr,
		})
	}

	server.ThrowException(&ExceptionResponse{
		StatusCode: http.StatusBadRequest,
		Success:    false,
		Message:    BadRequestMessage,
		Error:      err.Error(),
		Errors:     []validations.ValidationError{},
	})
}

func (server *HttpServer) ThrowInternalServerException(err error) {
	server.ThrowException(&ExceptionResponse{
		StatusCode: http.StatusInternalServerError,
		Success:    false,
		Message:    InternalServerMessage,
		Error:      err.Error(),
		Errors:     []validations.ValidationError{},
	})
}

func (server *HttpServer) ThrowForbiddenException(err error) {
	server.ThrowException(&ExceptionResponse{
		StatusCode: http.StatusForbidden,
		Success:    false,
		Message:    ForbiddenMessage,
		Error:      err.Error(),
		Errors:     []validations.ValidationError{},
	})
}

type DbException struct {
	Err             error
	NotFoundMessage string `default:"Not Found"`
}

func (server *HttpServer) ThrowDbException(param DbException) {
	typ := reflect.TypeOf(param)

	if param.Err == sql.ErrNoRows {
		if param.NotFoundMessage == "" {
			field, _ := typ.FieldByName("NotFoundMessage")
			param.NotFoundMessage = field.Tag.Get("default")
		}
		server.ThrowForbiddenException(fmt.Errorf(param.NotFoundMessage))
		return
	}

	server.ThrowInternalServerException(param.Err)
}
