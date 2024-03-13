package main

import (
	"database/sql"
	"os"

	"github.com/duyanhitbe/library-golang/apis"
	"github.com/duyanhitbe/library-golang/config"
	"github.com/duyanhitbe/library-golang/token"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	//Env
	env := config.NewEnv()

	//Database
	db, err := sql.Open(env.DriverName, env.DataSource)
	if err != nil {
		log.Fatal().Err(err)
	}
	errPing := db.Ping()
	if errPing != nil {
		log.Fatal().Err(errPing)
	}

	//Token maker
	tokenMaker := token.NewJWTMaker(env.SecretJWT)

	//Server
	server := apis.NewHttpServer(env, db, tokenMaker)

	//Start the server
	err = server.Start()
	if err != nil {
		log.Fatal().Err(err)
	}
}
