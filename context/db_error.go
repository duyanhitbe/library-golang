package context

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
)

type DbException struct {
	Ctx             *gin.Context
	Err             error
	NotFoundMessage string `default:"Not Found"`
}

func ThrowDbException(param DbException) {
	typ := reflect.TypeOf(param)

	if param.Err == sql.ErrNoRows {
		if param.NotFoundMessage == "" {
			field, _ := typ.FieldByName("NotFoundMessage")
			param.NotFoundMessage = field.Tag.Get("default")
		}

		exception := ForbiddenException(fmt.Errorf(param.NotFoundMessage))
		ThrowException(param.Ctx, exception)
		return
	}

	exception := InternalServerException(param.Err)
	ThrowException(param.Ctx, exception)
}
