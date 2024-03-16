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
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	os.Exit(m.Run())
}

func removeAll() {
	removeAllUser()
	removeAllBook()
	removeAllBookInfo()
	removeAllCategory()
	removeAllBorrower()
}

func removeAllCategory() {
	testDB.QueryRowContext(context.Background(), `DELETE FROM "categories"`)
}

func removeAllUser() {
	testDB.QueryRowContext(context.Background(), `DELETE FROM "users"`)
}

func removeAllBook() {
	testDB.QueryRowContext(context.Background(), `DELETE FROM "books"`)
}

func removeAllBookInfo() {
	testDB.QueryRowContext(context.Background(), `DELETE FROM "book_infos"`)
}

func removeAllBorrower() {
	testDB.QueryRowContext(context.Background(), `DELETE FROM "borrowers"`)
}
