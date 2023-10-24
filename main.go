package main

import (
	"context"
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

func createAdmin(transaction *db.Transaction) (err error) {
	var adminUser db.User
	ctx := context.Background()
	err = transaction.ExecTx(ctx, func(q *db.Queries) error {
		// 1. find admin if exists
		var queryErr error
		adminUser, queryErr = transaction.GetUserByUsernameForUpdate(ctx, config.AdminUsername)
		if queryErr != nil {
			if queryErr != sql.ErrNoRows {
				log.Fatal("error deleting admin: ", err)
			} else {
				// 2. create admin if not exists
				hashedPassword, hashErr := util.HashPassword(config.AdminPassword)
				if hashErr != nil {
					log.Fatal("error hashing admin password: ", err)
				}
				_, createErr := transaction.CreateUser(context.Background(), db.CreateUserParams{
					Username:    config.AdminUsername,
					Password:    hashedPassword,
					Description: sql.NullString{String: "ONE ACCOUNT TO RULE THEM ALL", Valid: true},
					Activated:   true,
				})
				if createErr != nil {
					log.Fatal("error creating admin: ", err)
				}
			}
		} else {
			// 2. or, update admin password if exists
			hashedPassword, hashErr := util.HashPassword(config.AdminPassword)
			if hashErr != nil {
				log.Fatal("error hashing admin password: ", err)
			}
			_, updateErr := transaction.UpdateUser(context.Background(), db.UpdateUserParams{
				Uid:         adminUser.Uid,
				Username:    config.AdminUsername,
				Password:    hashedPassword,
				Description: sql.NullString{String: "ONE ACCOUNT TO RULE THEM ALL", Valid: true},
				Activated:   true,
			})
			if updateErr != nil {
				log.Fatal("error updating admin: ", err)
			}
		}
		return nil
	})
	return
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

	// make sure admin is in database
	err = createAdmin(transaction)
	if err != nil {
		log.Fatal("cannot create admin: ", err)
	}

	server, err := api.NewServer(config, transaction)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start("0.0.0.0:8080")

	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
