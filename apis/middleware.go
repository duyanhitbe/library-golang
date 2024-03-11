package apis

import "github.com/gin-gonic/gin"

func (server *HttpServer) setCtx() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		server.ctx = ctx
		ctx.Next()
	}
}
