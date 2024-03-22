package apis

import (
	"github.com/duyanhitbe/library-golang/db"
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

// CreateUser godoc
// @Summary Create one user
// @Tags User API
// @Accept application/json
// @Produce application/json
// @Param body body apis.CreateUserRequest true "Create user request"
// @Security BearerAuth
// @Success 200 {object} apis.SuccessResponse{data=apis.UserResponse} "success"
// @Failure 400 {object} apis.ExceptionResponse "client error"
// @Failure 500 {object} apis.ExceptionResponse "database error"
// @Router /v1/user [post]
func (server *HttpServer) CreateUser(ctx *gin.Context) {
	var req CreateUserRequest
	if ok := server.BindJSON(&req); !ok {
		return
	}

	hashedPassword, err := server.hash.Hash(req.Password)
	if err != nil {
		server.ThrowInternalServerException(err)
		return
	}
	user, err := server.store.CreateUser(ctx, db.CreateUserParams{
		Username: req.Username,
		Password: hashedPassword,
		Role:     db.RoleEnum(req.Role),
	})
	if err != nil {
		server.ThrowInternalServerException(err)
		return
	}

	rsp := server.parseUserResponse(user)
	server.OkResponse(rsp)
}

// ListUser godoc
// @Summary Get a list of users
// @Tags User API
// @Accept application/json
// @Produce application/json
// @Param query query apis.SwaggerListRequest false "List query request"
// @Security BearerAuth
// @Success 200 {object} apis.PaginationResponse{data=[]apis.UserResponse} "success"
// @Failure 400 {object} apis.ExceptionResponse "client error"
// @Failure 500 {object} apis.ExceptionResponse "database error"
// @Router /v1/user [get]
func (server *HttpServer) ListUser(ctx *gin.Context) {
	req := server.BindPagination()
	if req == nil {
		return
	}

	params := db.ListUserParams{
		Limit:  int32(req.Limit),
		Offset: int32(req.Offset),
	}
	users, err := server.store.ListUser(ctx, params)
	if err != nil {
		server.ThrowDbException(DbException{
			Err: err,
		})
		return
	}
	total, err := server.store.CountUser(ctx)
	if err != nil {
		server.ThrowDbException(DbException{
			Err: err,
		})
		return
	}

	result := []*UserResponse{}
	for _, user := range users {
		rsp := server.parseUserResponse(user)
		result = append(result, rsp)
	}

	server.PaginatedResponse(req, result, total)
}

// GetOneUserById godoc
// @Summary Get one user by id
// @Tags User API
// @Accept application/json
// @Produce application/json
// @Param id path string true "user id"
// @Security BearerAuth
// @Success 200 {object} apis.SuccessResponse{data=apis.UserResponse} "success"
// @Failure 400 {object} apis.ExceptionResponse "client error"
// @Failure 500 {object} apis.ExceptionResponse "database error"
// @Router /v1/user/{id} [get]
func (server *HttpServer) GetOneUserById(ctx *gin.Context) {
	id, ok := server.BindID()
	if !ok {
		return
	}

	user, err := server.store.GetOneUserById(ctx, *id)
	if err != nil {
		server.ThrowDbException(DbException{
			Err:             err,
			NotFoundMessage: "User not found",
		})
		return
	}

	rsp := server.parseUserResponse(user)
	server.OkResponse(rsp)
}

type UpdateOneUserByIdRequest struct {
	Username string `json:"username" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

// UpdateOneUserById godoc
// @Summary Update one user by id
// @Tags User API
// @Accept application/json
// @Produce application/json
// @Param id path string true "user id"
// @Param body body apis.UpdateOneUserByIdRequest true "Update user request"
// @Security BearerAuth
// @Success 200 {object} apis.SuccessResponse{data=apis.UserResponse} "success"
// @Failure 400 {object} apis.ExceptionResponse "client error"
// @Failure 500 {object} apis.ExceptionResponse "database error"
// @Router /v1/user/{id} [patch]
func (server *HttpServer) UpdateOneUserById(ctx *gin.Context) {
	var req UpdateOneUserByIdRequest
	if ok := server.BindJSON(&req); !ok {
		return
	}
	id, ok := server.BindID()
	if !ok {
		return
	}

	params := db.UpdateOneUserByIdParams{
		ID:       *id,
		Username: req.Username,
		Role:     db.RoleEnum(req.Role),
	}
	user, err := server.store.UpdateOneUserById(ctx, params)
	if err != nil {
		server.ThrowDbException(DbException{
			Err:             err,
			NotFoundMessage: "User not found",
		})
		return
	}

	rsp := server.parseUserResponse(user)
	server.OkResponse(rsp)
}

// DeleteOneUserById godoc
// @Summary Delete one user by id
// @Tags User API
// @Accept application/json
// @Produce application/json
// @Param id path string true "user id"
// @Security BearerAuth
// @Success 200 {object} apis.SuccessResponse{data=apis.UserResponse} "success"
// @Failure 400 {object} apis.ExceptionResponse "client error"
// @Failure 500 {object} apis.ExceptionResponse "database error"
// @Router /v1/user/{id} [delete]
func (server *HttpServer) DeleteOneUserById(ctx *gin.Context) {
	id, ok := server.BindID()
	if !ok {
		return
	}

	user, err := server.store.DeleteOneUserById(ctx, *id)
	if err != nil {
		server.ThrowDbException(DbException{
			Err:             err,
			NotFoundMessage: "User not found",
		})
		return
	}

	rsp := server.parseUserResponse(user)
	server.OkResponse(rsp)
}
