package api

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testServer Server

func TestMain(m *testing.M) {
	var err error
	apitestDB, err := sql.Open("postgres", "postgresql://todo:todo@localhost:7899/todo_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	testServer.db = apitestDB

	os.Exit(m.Run())
}
