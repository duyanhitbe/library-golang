package api

import (
	"database/sql"
	"net/http"

	"github.com/duyanhitbe/library-golang/config"
	"github.com/duyanhitbe/library-golang/db"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	store  *db.Store
	env    *config.Env
}

func NewServer(env *config.Env, database *sql.DB) *Server {
	engine := gin.Default()
	store := db.NewStore(database)
	return &Server{
		engine: engine,
		store:  store,
		env:    env,
	}
}

// Start HTTP server
func (server *Server) Start() error {
	engine := server.engine
	serverAddress := server.env.ServerAddress

	mapUserRouter(engine)

	return server.engine.Run(serverAddress)
}

// Set up routing for user module
func mapUserRouter(engine *gin.Engine) {
	engine.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
}
