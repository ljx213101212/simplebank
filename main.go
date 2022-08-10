package main

import (
	"database/sql"
	"log"

	"github.com/ljx213101212/simplebank/api"
	db "github.com/ljx213101212/simplebank/db/sqlc"
	"github.com/techschool/simplebank/util"
)

func main() {

	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(store)

	if err != nil {
		log.Fatal("cannot create new server:", err)
	}
	err = server.Start(config.HTTPServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
