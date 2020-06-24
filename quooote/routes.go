package quooote

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Index (w http.ResponseWriter, r *http.Request) {
	render(w, "index.html", r)
}

func FetchQuotes (w http.ResponseWriter, r *http.Request) {
	render(w, "posted_quote.html", r)
}

func AddQuote (w http.ResponseWriter, r *http.Request) {
	render(w, "add_quote.html", r)
}

func PostQuote (w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	title := r.FormValue("quote_title")
	body := r.FormValue("quote_body")
	quotes:= Quote{
		Title:  title,
		Body:   body,
		Author: "Gabriel",
	}
	renderWithData(w, "posted_quote.html", r, quotes)
}

// Pending Quotes Handler Functions

func Pending(w http.ResponseWriter, r *http.Request) {
	data := getPendingQuotes()
	renderWithData(w, "pending_quotes.html", r, data)
}

func ViewQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	quooteToBeRendered := findQuote(id)
	renderWithData(w, "specific_quote.html", r, quooteToBeRendered)
}

func CreateQuote(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	quote := Quote{}
	quote.Author = 2
	quote.Title = r.FormValue("title")
	quote.Body = r.FormValue("body")
	createQuoteWith(&quote)
}

func createQuoteWith(data *Quote) {
	_ = data
}
