package context

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func getValidationError(err error) ([]ValidationError, bool) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ValidationError, len(ve))
		for i, fe := range ve {
			field := fe.Field()
			tag := fe.Tag()
			params := fe.Param()
			fieldType := fe.Type().String()
			out[i] = ValidationError{
				Field: field,
				Error: msgForTag(field, tag, params, fieldType),
			}
		}
		return out, true
	}
	return []ValidationError{}, false
}

func msgForTag(field, tag, params, fieldType string) string {
	switch tag {
	case "required":
		return field + " is required"
	case "email":
		return field + " is not an email"
	case "oneof":
		return field + " must be one of " + params
	case "min":
		if fieldType == "string" {
			return field + "'s length must greater than " + params
		} else {
			return field + "must greater than " + params
		}
	case "max":
		if fieldType == "string" {
			return field + "'s length must less than " + params
		} else {
			return field + "must less than " + params
		}
	}

	return ""
}

func ParseUUID(ctx *gin.Context, id string) *uuid.UUID {
	result, err := uuid.Parse(id)
	if err != nil {
		exception := BadRequestException(err)
		ThrowException(ctx, exception)
		return nil
	}
	return &result
}
