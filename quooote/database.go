package quooote

import (
	"database/sql"
	"log"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "53Z4YPjvtiKsE8FF"
	dbName := "quooote"

	db, err := sql.Open(dbDriver, dbUser + ":" + dbPass + "@/" + dbName)
	if err != nil {
		log.Fatal(err)
	}

	return db
}