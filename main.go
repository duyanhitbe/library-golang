package main

import (
	"database/sql"
	"os"

	"github.com/duyanhitbe/library-golang/apis"
	"github.com/duyanhitbe/library-golang/config"
	_ "github.com/duyanhitbe/library-golang/docs"
	"github.com/duyanhitbe/library-golang/token"
	_ "github.com/go-openapi/jsonpointer"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// @title Library App
// @version 1.0
// @description Library App using Golang
// @host localhost:8080
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	//Env
	env := config.NewEnv()

	//Database
	db, err := connectDB(env)
	if err != nil {
		log.Fatal().Err(err).Msg("fail to connect database")
	}

	//Migrate
	err = runMigration(env.DataSource)
	if err != nil {
		log.Fatal().Err(err).Msg("fail to run migration")
	}

	//Start HTTP server
	err = startHttpServer(db, env)
	if err != nil {
		log.Fatal().Err(err).Msg("fail to start HTTP server")
	}
}

func connectDB(env *config.Env) (*sql.DB, error) {
	db, err := sql.Open(env.DriverName, env.DataSource)
	if err != nil {
		return nil, err
	}
	errPing := db.Ping()
	if errPing != nil {
		return nil, err
	}
	return db, nil
}

func runMigration(dataSource string) error {
	m, err := migrate.New("file://sql/migrations", dataSource)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	log.Info().Msg("Run migration successfully")
	return nil
}

func startHttpServer(db *sql.DB, env *config.Env) error {
	//Token maker
	tokenMaker := token.NewJWTMaker(env.SecretJWT)

	//Server
	server := apis.NewHttpServer(env, db, tokenMaker)

	//Start the server
	log.Info().Msg("Server start successfully")
	log.Info().Msgf("Swagger: http://%s/docs/index.html", env.HttpServerAddress)
	return server.Start()
}
