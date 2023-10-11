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
var dbSecret string

var testQueries *Queries

func init() {
	secretConfig := util.SecretConfig{}

	err := util.LoadConfig(&secretConfig, "../../")
	if err != nil {
		log.Fatal("cannot load secret config: ", err)
	}

	dbSecret = secretConfig.PostgresPassword
}

func TestMain(m *testing.M) {
	var dbDriver = "postgres"
	var dbSource = "postgresql://postgres:" + dbSecret + "@localhost:5432/inventory_management_server_db?sslmode=disable"

	// connect to database
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
