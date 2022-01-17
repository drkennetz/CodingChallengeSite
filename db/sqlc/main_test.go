package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	_ "github.com/lib/pq"
)

const (
	// should source these from environment in future
	dbDriver = "postgres"
	dbSource = "postgresql://root:harrison40@localhost:5432/coding_challenge?sslmode=disable"
)

var testQueries *Queries
// global test db conn
var testDB *sql.DB

func TestMain(m *testing.M) {
	// main entry point of all tests of a specific golang package
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}