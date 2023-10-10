package db

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

// read secret from environment variables
var dbSecret string
var dbDriver = "postgres"
var dbSource = "postgresql://root:" + dbSecret + "@localhost:5432/simple_bank?sslmode=disable"

var testQueries *Queries

func init() {
	dbSecret = os.Getenv("POSTGRES_PASSWORD")
}

func TestMain(m *testing.M) {
	// connect to database
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
