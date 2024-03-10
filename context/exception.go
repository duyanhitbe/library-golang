package context

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	BadRequestMessage     = "Bad Request"
	InternalServerMessage = "Internal Server"
	ForbiddenMessage      = "Forbidden"
)

type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type ExceptionResponse struct {
	StatusCode int               `json:"status_code"`
	Success    bool              `json:"success"`
	Message    string            `json:"message"`
	Error      string            `json:"error"`
	Errors     []ValidationError `json:"errors"`
}

func ThrowException(ctx *gin.Context, exception *ExceptionResponse) {
	ctx.JSON(exception.StatusCode, exception)
}

func BadRequestException(err error) *ExceptionResponse {
	if validationErr, isValidationErr := getValidationError(err); isValidationErr {
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
		Errors:     []ValidationError{},
	}
}

func InternalServerException(err error) *ExceptionResponse {
	return &ExceptionResponse{
		StatusCode: http.StatusInternalServerError,
		Success:    false,
		Message:    InternalServerMessage,
		Error:      err.Error(),
		Errors:     []ValidationError{},
	}
}

func ForbiddenException(err error) *ExceptionResponse {
	return &ExceptionResponse{
		StatusCode: http.StatusForbidden,
		Success:    false,
		Message:    ForbiddenMessage,
		Error:      err.Error(),
		Errors:     []ValidationError{},
	}
}
