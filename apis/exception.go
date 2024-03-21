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
	UnauthorizedMessage   = "Unauthorized"
)

type ExceptionResponse struct {
	StatusCode int                           `json:"status_code,omitempty"`
	Success    bool                          `json:"success,omitempty"`
	Message    string                        `json:"message,omitempty"`
	Error      string                        `json:"error,omitempty"`
	Errors     []validations.ValidationError `json:"errors,omitempty"`
}

func (server *HttpServer) ThrowException(exception *ExceptionResponse) {
	server.ctx.JSON(exception.StatusCode, exception)
}

func NewBadRequestException(err error) *ExceptionResponse {
	if validationErr, isValidationErr := validations.GetValidationError(err); isValidationErr {
		return &ExceptionResponse{
			StatusCode: http.StatusBadRequest,
			Success:    false,
			Message:    BadRequestMessage,
			Error:      "Validation Error",
			Errors:     validationErr,
		}
	}

	return &ExceptionResponse{
		StatusCode: http.StatusBadRequest,
		Success:    false,
		Message:    BadRequestMessage,
		Error:      err.Error(),
		Errors:     []validations.ValidationError{},
	}
}

func (server *HttpServer) ThrowBadRequestException(err error) {
	server.ThrowException(NewBadRequestException(err))
}

func NewInternalServerException(err error) *ExceptionResponse {
	return &ExceptionResponse{
		StatusCode: http.StatusInternalServerError,
		Success:    false,
		Message:    InternalServerMessage,
		Error:      err.Error(),
		Errors:     []validations.ValidationError{},
	}
}

func (server *HttpServer) ThrowInternalServerException(err error) {
	server.ThrowException(NewInternalServerException(err))
}

func NewForbiddenException(err error) *ExceptionResponse {
	return &ExceptionResponse{
		StatusCode: http.StatusForbidden,
		Success:    false,
		Message:    ForbiddenMessage,
		Error:      err.Error(),
		Errors:     []validations.ValidationError{},
	}
}

func (server *HttpServer) ThrowForbiddenException(err error) {
	server.ThrowException(NewForbiddenException(err))
}

func NewUnauthorizedException(err error) *ExceptionResponse {
	return &ExceptionResponse{
		StatusCode: http.StatusUnauthorized,
		Success:    false,
		Message:    UnauthorizedMessage,
		Error:      err.Error(),
		Errors:     []validations.ValidationError{},
	}
}

func (server *HttpServer) ThrowUnauthorizedException(err error) {
	server.ThrowException(NewUnauthorizedException(err))
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
