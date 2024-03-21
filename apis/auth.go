package apis

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginUser godoc
// @Summary Login User
// @Tags Auth API
// @Accept application/json
// @Produce application/json
// @Param body body apis.LoginUserRequest true "Login user request"
// @Success 200 {object} apis.SuccessResponse{data=apis.LoginUserResponse} "success"
// @Failure 400 {object} apis.ExceptionResponse "client error"
// @Failure 500 {object} apis.ExceptionResponse "database error"
// @Router /v1/auth/user/login [post]
func (server *HttpServer) LoginUser(ctx *gin.Context) {
	var req LoginUserRequest
	if ok := server.BindJSON(&req); !ok {
		return
	}

	//Find user
	user, err := server.store.GetOneUserByUsername(ctx, req.Username)
	if err != nil {
		server.ThrowDbException(DbException{
			Err:             err,
			NotFoundMessage: "User not found",
		})
		return
	}

	//Verify password
	ok, err := server.hash.Verify(user.Password, req.Password)
	if err != nil {
		server.ThrowInternalServerException(err)
		return
	}
	if !ok {
		server.ThrowUnauthorizedException(errors.New("Wrong password"))
		return
	}

	userId := user.ID
	accessTokenDuration := time.Hour
	accessToken, accessTokenPayload, err := server.tokenMaker.Generate(userId, user.Role, accessTokenDuration)
	if err != nil {
		server.ThrowInternalServerException(err)
		return
	}

	rsp := LoginUserResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessTokenPayload.ExpiresAt,
		AccessTokenIssuedAt:  accessTokenPayload.IssuedAt,
	}
	server.OkResponse(rsp)
}
