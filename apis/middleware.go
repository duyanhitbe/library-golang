package apis

import (
	"errors"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/duyanhitbe/library-golang/db"
	"github.com/duyanhitbe/library-golang/token"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

const (
	authorizationHeaderKey = "authorization"
	authorizationType      = "bearer"
	authorizationPayload   = "authorizationPayload"
	userId                 = "userId"
	userRole               = "userRole"
)

var (
	missingAuthorizationError   = errors.New("missing authorization")
	malformedAuthorizationError = errors.New("malformed authorization")
	invalidAuthorizationError   = errors.New("invalid authorization")
	permissionError             = errors.New("permission deny")
)

func (server *HttpServer) setCtx() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		server.ctx = ctx
		ctx.Next()
	}
}

func logger(ctx *gin.Context) {
	start := time.Now()
	ctx.Next()
	duration := time.Since(start)
	logger := log.Info()

	logger.Str("protocol", "http").
		Str("method", ctx.Request.Method).
		Str("path", ctx.Request.RequestURI).
		Dur("duration", duration).
		Msg("receive a HTTP request")
}

func authenticateMiddleware(tokenMaker token.TokenMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			exception := NewUnauthorizedException(missingAuthorizationError)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, exception)
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) != 2 {
			exception := NewUnauthorizedException(malformedAuthorizationError)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, exception)
			return
		}

		prefix := strings.ToLower(fields[0])
		if prefix != authorizationType {
			exception := NewUnauthorizedException(invalidAuthorizationError)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, exception)
			return
		}

		token := fields[1]
		payload, err := tokenMaker.Verify(token)
		if err != nil {
			exception := NewUnauthorizedException(invalidAuthorizationError)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, exception)
			return
		}
		ctx.Set(authorizationPayload, payload)
		ctx.Set(userId, payload.UserID)
		ctx.Set(userRole, payload.UserRole)

		ctx.Next()
	}
}

func authorizeMiddleware(roles []db.RoleEnum) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roleString, ok := ctx.Get(userRole)
		if !ok {
			exception := NewUnauthorizedException(invalidAuthorizationError)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, exception)
			return
		}

		role, ok := roleString.(db.RoleEnum)
		if !ok {
			err := errors.New("Can not parse user role")
			exception := NewInternalServerException(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, exception)
			return
		}

		if !slices.Contains(roles, role) {
			exception := NewForbiddenException(permissionError)
			ctx.AbortWithStatusJSON(http.StatusForbidden, exception)
			return
		}

		ctx.Next()
	}
}
