package quooote

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "monk"
	dbPass := "53Z4YPjvtiKsE8FF"
	dbName := "quooote"

	db, err := sql.Open(dbDriver, dbUser + ":" + dbPass + "@/" + dbName)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func getPendingQuotes() (pendingQuotes []Quoote){
	db := dbConn()

	PendingQuote := Quoote{}
	PendingQuotes := []Quoote{}
	stmtOut, err := db.Query("SELECT * FROM quotes WHERE status = ?", 1)
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

		PendingQuote.Id = id
		PendingQuote.Body = body
		PendingQuote.Punchline = punchline
		PendingQuote.StartedBy = startedBy
		PendingQuote.Status = status

		PendingQuotes = append(PendingQuotes, PendingQuote)
	}
	return PendingQuotes
}