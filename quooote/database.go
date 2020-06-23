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

func (q *Quoote) getPendingQuotes() (pendingQuotes []Quoote){
	db := dbConn()
	PendingQuotes := []Quoote{}
	stmtOut, err := db.Query("SELECT * FROM quootes WHERE status = ?", 1)
	if err != nil {
		log.Fatal(err)
	}

	for stmtOut.Next() {
		var id int
		var punchline string
		var body string
		var startedBy int
		var status int


		_ = stmtOut.Scan(&id, &punchline, &body, &startedBy, &status)

		q.Id = id
		q.Body = body
		q.Punchline = punchline
		q.StartedBy = startedBy
		q.Status = status

		PendingQuotes = append(PendingQuotes, q)
	}
	return PendingQuotes
}