package main

import (
	"database/sql"
	"log"

	"github.com/Peiyang-Aeromodelling-Association/inventory_management_server/api"
	db "github.com/Peiyang-Aeromodelling-Association/inventory_management_server/db/sqlc"
	"github.com/Peiyang-Aeromodelling-Association/inventory_management_server/util"
	_ "github.com/lib/pq"
)

// read secret from environment variables
var config util.Config

func init() {
	var err error
	config, err = util.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load secret config: ", err)
	}
}

func main() {
	var dbDriver = "postgres"
	var dbSource = "postgresql://postgres:" + config.PostgresPassword + "@localhost:5432/inventory_management_server_db?sslmode=disable"

	// connect to database
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	transaction := db.NewTransaction(conn)
	server, err := api.NewServer(config, transaction)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start("0.0.0.0:8080")

	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
