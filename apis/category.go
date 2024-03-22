package apis

import (
	"github.com/duyanhitbe/library-golang/db"
	"github.com/gin-gonic/gin"
)

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

// CreateCategory godoc
// @Summary Create one category
// @Tags Category API
// @Accept application/json
// @Produce application/json
// @Param body body apis.CreateCategoryRequest true "Create category request"
// @Security BearerAuth
// @Success 200 {object} apis.SuccessResponse{data=db.Category} "success"
// @Failure 400 {object} apis.ExceptionResponse "client error"
// @Failure 500 {object} apis.ExceptionResponse "database error"
// @Router /v1/category [post]
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

// ListCategory godoc
// @Summary Get a list of categories
// @Tags Category API
// @Accept application/json
// @Produce application/json
// @Param query query apis.SwaggerListRequest false "List query request"
// @Success 200 {object} apis.PaginationResponse{data=[]db.Category} "success"
// @Failure 400 {object} apis.ExceptionResponse "client error"
// @Failure 500 {object} apis.ExceptionResponse "database error"
// @Router /v1/category [get]
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

// GetOneCategoryById godoc
// @Summary Get one category by id
// @Tags Category API
// @Accept application/json
// @Produce application/json
// @Param id path string true "category id"
// @Success 200 {object} apis.SuccessResponse{data=db.Category} "success"
// @Failure 400 {object} apis.ExceptionResponse "client error"
// @Failure 500 {object} apis.ExceptionResponse "database error"
// @Router /v1/category/{id} [get]
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

// UpdateOneCategoryById godoc
// @Summary Update one category by id
// @Tags Category API
// @Accept application/json
// @Produce application/json
// @Param id path string true "category id"
// @Param body body apis.UpdateOneCategoryByIdRequest true "Update category request"
// @Security BearerAuth
// @Success 200 {object} apis.SuccessResponse{data=db.Category} "success"
// @Failure 400 {object} apis.ExceptionResponse "client error"
// @Failure 500 {object} apis.ExceptionResponse "database error"
// @Router /v1/category/{id} [patch]
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

// DeleteOneCategoryById godoc
// @Summary Delete one category by id
// @Tags Category API
// @Accept application/json
// @Produce application/json
// @Param id path string true "category id"
// @Security BearerAuth
// @Success 200 {object} apis.SuccessResponse{data=db.Category} "success"
// @Failure 400 {object} apis.ExceptionResponse "client error"
// @Failure 500 {object} apis.ExceptionResponse "database error"
// @Router /v1/category/{id} [delete]
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
