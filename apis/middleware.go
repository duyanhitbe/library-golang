package apis

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
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
