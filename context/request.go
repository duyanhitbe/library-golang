package context

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ListRequest struct {
	Limit  int64 `form:"limit"`
	Page   int64 `form:"page"`
	Offset int64
}

func BindPagination(ctx *gin.Context) *ListRequest {
	var req ListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		exception := BadRequestException(err)
		ThrowException(ctx, exception)
		return nil
	}
	if req.Limit == 0 {
		req.Limit = 10
	}
	if req.Page == 0 {
		req.Page = 1
	}
	req.Offset = req.Limit * (req.Page - 1)
	return &req
}

type IdRequest struct {
	Id string `uri:"id" binding:"required"`
}

func BindID(ctx *gin.Context) (*uuid.UUID, bool) {
	var req IdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		exception := BadRequestException(err)
		ThrowException(ctx, exception)
		return nil, false
	}

	id := ParseUUID(ctx, req.Id)

	return id, id != nil
}

func BindJSON(ctx *gin.Context, req interface{}) bool {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		exception := BadRequestException(err)
		ThrowException(ctx, exception)
		return false
	}
	return true
}
