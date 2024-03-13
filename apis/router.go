package apis

import (
	"net/http"

	"github.com/duyanhitbe/library-golang/db"
	"github.com/gin-gonic/gin"
)

var (
	adminPermissionAuthorize   = authorizeMiddleware([]db.RoleEnum{db.RoleEnumADMIN})
	managerPermissionAuthorize = authorizeMiddleware([]db.RoleEnum{db.RoleEnumMANAGER})
	fullPermissionAuthorize    = authorizeMiddleware([]db.RoleEnum{db.RoleEnumADMIN, db.RoleEnumMANAGER})
)

func (server *HttpServer) initRouter() {
	//Middleware
	authenticate := authenticateMiddleware(server.tokenMaker)
	server.engine.Use(server.setCtx())
	server.engine.Use(logger)

	//Init routers
	server.engine.GET("/", homePath)
	server.initCategoryRouter(&authenticate)
	server.initUserRouter(&authenticate)
	server.initBookRouter(&authenticate)
	server.initBorrowerRouter(&authenticate)
	server.initAuthRouter(&authenticate)
}

func homePath(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Library",
	})
}

func (server *HttpServer) initCategoryRouter(authenticate *gin.HandlerFunc) {
	v1 := server.engine.Group("/v1/category")
	{
		{
			v1.POST("/", *authenticate, fullPermissionAuthorize, server.CreateCategory)
			v1.GET("/", server.ListCategory)
			v1.GET("/:id", server.GetOneCategoryById)
			v1.PATCH("/:id", *authenticate, fullPermissionAuthorize, server.UpdateOneCategoryById)
			v1.DELETE("/:id", *authenticate, fullPermissionAuthorize, server.DeleteOneCategoryById)
		}
	}
}

func (server *HttpServer) initUserRouter(authenticate *gin.HandlerFunc) {
	v1 := server.engine.Group("/v1/user")
	v1.Use(*authenticate)
	v1.Use(adminPermissionAuthorize)
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

func (server *HttpServer) initBookRouter(authenticate *gin.HandlerFunc) {
	v1 := server.engine.Group("/v1/book")
	{
		{
			v1.POST("/", *authenticate, fullPermissionAuthorize, server.CreateBook)
			v1.POST("/borrow", *authenticate, fullPermissionAuthorize, server.BorrowBook)
			v1.GET("/borrow/:id", *authenticate, fullPermissionAuthorize, server.ListBookByBorrowerId)
			v1.GET("/", server.ListBook)
			v1.GET("/:id", server.GetOneBookById)
			v1.PATCH("/:id", *authenticate, fullPermissionAuthorize, server.UpdateOneBookById)
			v1.DELETE("/:id", *authenticate, fullPermissionAuthorize, server.DeleteOneBookById)
		}
	}
}

func (server *HttpServer) initBorrowerRouter(authenticate *gin.HandlerFunc) {
	v1 := server.engine.Group("/v1/borrower")
	{
		{
			v1.GET("/book/:id", *authenticate, fullPermissionAuthorize, server.ListBorrowerByBookId)
		}
	}
}

func (server *HttpServer) initAuthRouter(authenticate *gin.HandlerFunc) {
	v1 := server.engine.Group("/v1/auth")
	{
		{
			v1.POST("/user/login", server.LoginUser)
		}
	}
}
