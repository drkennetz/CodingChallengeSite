package main

import (
	"database/sql"
	"log"

	"github.com/drkennetz/codingchallenge/api"
	db "github.com/drkennetz/codingchallenge/db/sqlc"
	"github.com/drkennetz/codingchallenge/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configurations.", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}

}