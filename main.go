package main

import (
	"database/sql"
	"log"

	"github.com/ljx213101212/simplebank/api"
	db "github.com/ljx213101212/simplebank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(store)

	if err != nil {
		log.Fatal("cannot create new server:", err)
	}
	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
