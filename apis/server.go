package apis

import (
	"database/sql"
	"net/http"

	"github.com/duyanhitbe/library-golang/config"
	"github.com/duyanhitbe/library-golang/db"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	engine *gin.Engine
	store  db.Store
	env    *config.Env
}

func NewHttpServer(env *config.Env, database *sql.DB) *HttpServer {
	engine := gin.Default()
	store := db.NewStore(database)
	return &HttpServer{
		engine: engine,
		store:  store,
		env:    env,
	}
}

// Start HTTP server
func (server *HttpServer) Start() error {
	serverAddress := server.env.HttpServerAddress

	server.setUpHomeRoute()
	server.setUpCategoryRoutes()

	return server.engine.Run(serverAddress)
}

func (server *HttpServer) setUpHomeRoute() {
	server.engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Library",
		})
	})
}

func (server *HttpServer) setUpCategoryRoutes() {
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
