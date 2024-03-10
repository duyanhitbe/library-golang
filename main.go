package main

import (
	"database/sql"
	"log"

	"github.com/duyanhitbe/library-golang/apis"
	"github.com/duyanhitbe/library-golang/config"
	_ "github.com/lib/pq"
)

func main() {
	env := config.NewEnv()

	db, err := sql.Open(env.DriverName, env.DataSource)
	if err != nil {
		log.Fatal(err)
	}
	errPing := db.Ping()
	if errPing != nil {
		log.Fatal(errPing)
	}

	server := apis.NewHttpServer(env, db)

	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
