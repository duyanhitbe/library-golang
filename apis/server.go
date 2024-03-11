package apis

import (
	"database/sql"

	"github.com/duyanhitbe/library-golang/config"
	"github.com/duyanhitbe/library-golang/db"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	engine *gin.Engine
	store  db.Store
	env    *config.Env
	ctx    *gin.Context
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

	server.initRouter()

	return server.engine.Run(serverAddress)
}
