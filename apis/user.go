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
