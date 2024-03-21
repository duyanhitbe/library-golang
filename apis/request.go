package apis

import (
	"github.com/duyanhitbe/library-golang/validations"
	"github.com/google/uuid"
)

type SwaggerListRequest struct {
	Limit int64 `form:"limit"`
	Page  int64 `form:"page"`
}

type ListRequest struct {
	Limit  int64 `form:"limit"`
	Page   int64 `form:"page"`
	Offset int64
}

func (server *HttpServer) BindPagination() *ListRequest {
	var req ListRequest
	if err := server.ctx.ShouldBindQuery(&req); err != nil {
		server.ThrowBadRequestException(err)
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

func (server *HttpServer) BindID() (*uuid.UUID, bool) {
	var req IdRequest
	if err := server.ctx.ShouldBindUri(&req); err != nil {
		server.ThrowBadRequestException(err)
		return nil, false
	}

	id, err := validations.ParseUUID(req.Id)
	if err != nil {
		server.ThrowBadRequestException(err)
		return nil, false
	}

	return id, id != nil
}

func (server *HttpServer) BindJSON(req interface{}) bool {
	if err := server.ctx.ShouldBindJSON(&req); err != nil {
		server.ThrowBadRequestException(err)
		return false
	}
	return true
}
