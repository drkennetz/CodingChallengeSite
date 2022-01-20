package main

import (
	"database/sql"
	"log"

	_"github.com/lib/pq"
	"github.com/drkennetz/codingchallenge/api"
	db "github.com/drkennetz/codingchallenge/db/sqlc"
)

const (
	// should source these from environment in future
	dbDriver = "postgres"
	dbSource = "postgresql://root:harrison40@localhost:5432/coding_challenge?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}

}