package apis

import (
	"github.com/duyanhitbe/library-golang/context"
	"github.com/duyanhitbe/library-golang/db"
	"github.com/gin-gonic/gin"
)

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *HttpServer) CreateCategory(ctx *gin.Context) {
	var req CreateCategoryRequest
	if ok := context.BindJSON(ctx, &req); !ok {
		return
	}

	id, err := server.store.CreateCategory(ctx, req.Name)
	if err != nil {
		exception := context.InternalServerException(err)
		context.ThrowException(ctx, exception)
		return
	}

	response := context.OkResponse(id)
	context.Response(ctx, response)
}

func (server *HttpServer) ListCategory(ctx *gin.Context) {
	req := context.BindPagination(ctx)
	if req == nil {
		return
	}

	params := db.ListCategoryParams{
		Limit:  int32(req.Limit),
		Offset: int32(req.Offset),
	}
	categories, err := server.store.ListCategory(ctx, params)
	if err != nil {
		context.ThrowDbException(context.DbException{
			Ctx: ctx,
			Err: err,
		})
		return
	}
	total, err := server.store.CountCategory(ctx)
	if err != nil {
		context.ThrowDbException(context.DbException{
			Ctx: ctx,
			Err: err,
		})
		return
	}

	response := context.PaginatedResponse(req, categories, total)
	context.Response(ctx, response)
}

func (server *HttpServer) GetOneCategoryById(ctx *gin.Context) {
	id, ok := context.BindID(ctx)
	if !ok {
		return
	}

	category, err := server.store.GetOneCategoryById(ctx, *id)
	if err != nil {
		context.ThrowDbException(context.DbException{
			Ctx:             ctx,
			Err:             err,
			NotFoundMessage: "Category not found",
		})
		return
	}

	response := context.OkResponse(category)
	context.Response(ctx, response)
}

type UpdateOneCategoryByIdRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *HttpServer) UpdateOneCategoryById(ctx *gin.Context) {
	var req UpdateOneCategoryByIdRequest
	if ok := context.BindJSON(ctx, &req); !ok {
		return
	}
	id, ok := context.BindID(ctx)
	if !ok {
		return
	}

	params := db.UpdateOneCategoryByIdParams{
		ID:   *id,
		Name: req.Name,
	}
	category, err := server.store.UpdateOneCategoryById(ctx, params)
	if err != nil {
		context.ThrowDbException(context.DbException{
			Ctx:             ctx,
			Err:             err,
			NotFoundMessage: "Category not found",
		})
		return
	}

	response := context.OkResponse(category)
	context.Response(ctx, response)
}

func (server *HttpServer) DeleteOneCategoryById(ctx *gin.Context) {
	id, ok := context.BindID(ctx)
	if !ok {
		return
	}

	category, err := server.store.DeleteOneCategoryById(ctx, *id)
	if err != nil {
		context.ThrowDbException(context.DbException{
			Ctx:             ctx,
			Err:             err,
			NotFoundMessage: "Category not found",
		})
		return
	}

	response := context.OkResponse(category)
	context.Response(ctx, response)
}
