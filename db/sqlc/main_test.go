package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Peiyang-Aeromodelling-Association/inventory_management_server/util"

	_ "github.com/lib/pq"
)

// read secret from environment variables
var config util.Config

var testQueries *Queries
var testDB *sql.DB

func init() {
	var err error
	config, err = util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load secret config: ", err)
	}
}

func TestMain(m *testing.M) {
	var err error

	var dbDriver = "postgres"
	var dbSource = "postgresql://postgres:" + config.PostgresPassword + "@localhost:5432/inventory_management_server_db?sslmode=disable"

	// connect to database
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
