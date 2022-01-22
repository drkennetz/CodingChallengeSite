package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/drkennetz/codingchallenge/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
// global test db conn
var testDB *sql.DB

func TestMain(m *testing.M) {
	// main entry point of all tests of a specific golang package
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("failed to load config.", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}