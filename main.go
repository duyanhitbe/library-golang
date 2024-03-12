package main

import (
	"database/sql"
	"os"

	"github.com/duyanhitbe/library-golang/apis"
	"github.com/duyanhitbe/library-golang/config"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	env := config.NewEnv()

	db, err := sql.Open(env.DriverName, env.DataSource)
	if err != nil {
		log.Fatal().Err(err)
	}
	errPing := db.Ping()
	if errPing != nil {
		log.Fatal().Err(errPing)
	}

	server := apis.NewHttpServer(env, db)

	err = server.Start()
	if err != nil {
		log.Fatal().Err(err)
	}
}
