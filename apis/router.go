package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *HttpServer) initRouter() {
	server.engine.Use(server.setCtx())
	server.engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Library",
		})
	})
	v1 := server.engine.Group("/v1")
	{
		categoryV1 := v1.Group("/category")
		{
			categoryV1.POST("/", server.CreateCategory)
			categoryV1.GET("/", server.ListCategory)
			categoryV1.GET("/:id", server.GetOneCategoryById)
			categoryV1.PATCH("/:id", server.UpdateOneCategoryById)
			categoryV1.DELETE("/:id", server.DeleteOneCategoryById)
		}
	}
}
