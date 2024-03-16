package db

import (
	"context"
	"database/sql"
	"github.com/duyanhitbe/library-golang/config"
	"github.com/rs/zerolog/log"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	env := config.NewEnv()

	db, err := sql.Open(env.DriverName, env.DataSource)
	if err != nil {
		log.Fatal().Msg("Can not connect DB")
	}

	testDB = db
	testQueries = New(testDB)

	os.Exit(m.Run())
}

func removeAllCategory() {
	testDB.QueryRowContext(context.Background(), `DELETE FROM "categories"`)
}
