package apis

import (
	"database/sql"

	"github.com/duyanhitbe/library-golang/config"
	"github.com/duyanhitbe/library-golang/db"
	"github.com/duyanhitbe/library-golang/hash"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	engine *gin.Engine
	store  db.Store
	env    *config.Env
	ctx    *gin.Context
	hash   hash.Hash
}

func NewHttpServer(env *config.Env, database *sql.DB) *HttpServer {
	engine := gin.Default()
	store := db.NewStore(database)
	hash := hash.NewArgon2()
	return &HttpServer{
		engine: engine,
		store:  store,
		env:    env,
		hash:   hash,
	}
}

// Start HTTP server
func (server *HttpServer) Start() error {
	serverAddress := server.env.HttpServerAddress

	server.initRouter()

	return server.engine.Run(serverAddress)
}
