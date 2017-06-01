package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// DB gets a db connection
var DB = getConnection()

func getConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:secret@localhost/account-confirmation?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
