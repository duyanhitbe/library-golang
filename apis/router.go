package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *HttpServer) initRouter() {
	server.engine.Use(server.setCtx())
	server.engine.Use(logger)
	server.engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Library",
		})
	})
	server.initCategoryRouter()
	server.initUserRouter()
	server.initBookRouter()
	server.initBorrowerRouter()
}

func (server *HttpServer) initCategoryRouter() {
	v1 := server.engine.Group("/v1/category")
	{
		{
			v1.POST("/", server.CreateCategory)
			v1.GET("/", server.ListCategory)
			v1.GET("/:id", server.GetOneCategoryById)
			v1.PATCH("/:id", server.UpdateOneCategoryById)
			v1.DELETE("/:id", server.DeleteOneCategoryById)
		}
	}
}

func (server *HttpServer) initUserRouter() {
	v1 := server.engine.Group("/v1/user")
	{
		{
			v1.POST("/", server.CreateUser)
			v1.GET("/", server.ListUser)
			v1.GET("/:id", server.GetOneUserById)
			v1.PATCH("/:id", server.UpdateOneUserById)
			v1.DELETE("/:id", server.DeleteOneUserById)
		}
	}
}

func (server *HttpServer) initBookRouter() {
	v1 := server.engine.Group("/v1/book")
	{
		{
			v1.POST("/", server.CreateBook)
			v1.POST("/borrow", server.BorrowBook)
			v1.GET("/borrow/:id", server.ListBookByBorrowerId)
			v1.GET("/", server.ListBook)
			v1.GET("/:id", server.GetOneBookById)
			v1.PATCH("/:id", server.UpdateOneBookById)
			v1.DELETE("/:id", server.DeleteOneBookById)
		}
	}
}

func (server *HttpServer) initBorrowerRouter() {
	v1 := server.engine.Group("/v1/borrower")
	{
		{
			v1.GET("/book/:id", server.ListBorrowerByBookId)
		}
	}
}
