package main

import (
	"database/sql"
	"log"

	"github.com/duyanhitbe/library-golang/api"
	"github.com/duyanhitbe/library-golang/config"
	_ "github.com/lib/pq"
)

func main() {
	env := config.NewEnv()

	db, err := sql.Open(env.DriverName, env.DataSource)
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(env, db)

	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
