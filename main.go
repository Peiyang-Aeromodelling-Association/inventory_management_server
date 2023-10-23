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
var dbSecret string

func init() {
	secretConfig := util.SecretConfig{}

	err := util.LoadConfig(&secretConfig, "./")
	if err != nil {
		log.Fatal("cannot load secret config: ", err)
	}

	dbSecret = secretConfig.PostgresPassword
}

func main() {
	var dbDriver = "postgres"
	var dbSource = "postgresql://postgres:" + dbSecret + "@localhost:5432/inventory_management_server_db?sslmode=disable"

	// connect to database
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	transaction := db.NewTransaction(conn)
	server := api.NewServer(transaction)

	err = server.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
