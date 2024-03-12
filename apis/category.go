package apis

import (
	"github.com/duyanhitbe/library-golang/db"
	"github.com/gin-gonic/gin"
)

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *HttpServer) CreateCategory(ctx *gin.Context) {
	var req CreateCategoryRequest
	if ok := server.BindJSON(&req); !ok {
		return
	}

	category, err := server.store.CreateCategory(ctx, req.Name)
	if err != nil {
		server.ThrowInternalServerException(err)
		return
	}

	server.OkResponse(category)
}

func (server *HttpServer) ListCategory(ctx *gin.Context) {
	req := server.BindPagination()
	if req == nil {
		return
	}

	params := db.ListCategoryParams{
		Limit:  int32(req.Limit),
		Offset: int32(req.Offset),
	}
	categories, err := server.store.ListCategory(ctx, params)
	if err != nil {
		server.ThrowDbException(DbException{
			Err: err,
		})
		return
	}
	total, err := server.store.CountCategory(ctx)
	if err != nil {
		server.ThrowDbException(DbException{
			Err: err,
		})
		return
	}

	server.PaginatedResponse(req, categories, total)
}

func (server *HttpServer) GetOneCategoryById(ctx *gin.Context) {
	id, ok := server.BindID()
	if !ok {
		return
	}

	category, err := server.store.GetOneCategoryById(ctx, *id)
	if err != nil {
		server.ThrowDbException(DbException{
			Err:             err,
			NotFoundMessage: "Category not found",
		})
		return
	}

	server.OkResponse(category)
}

type UpdateOneCategoryByIdRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *HttpServer) UpdateOneCategoryById(ctx *gin.Context) {
	var req UpdateOneCategoryByIdRequest
	if ok := server.BindJSON(&req); !ok {
		return
	}
	id, ok := server.BindID()
	if !ok {
		return
	}

	params := db.UpdateOneCategoryByIdParams{
		ID:   *id,
		Name: req.Name,
	}
	category, err := server.store.UpdateOneCategoryById(ctx, params)
	if err != nil {
		server.ThrowDbException(DbException{
			Err:             err,
			NotFoundMessage: "Category not found",
		})
		return
	}

	server.OkResponse(category)
}

func (server *HttpServer) DeleteOneCategoryById(ctx *gin.Context) {
	id, ok := server.BindID()
	if !ok {
		return
	}

	category, err := server.store.DeleteOneCategoryById(ctx, *id)
	if err != nil {
		server.ThrowDbException(DbException{
			Err:             err,
			NotFoundMessage: "Category not found",
		})
		return
	}

	server.OkResponse(category)
}
