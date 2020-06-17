package quooote

import (
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
	quotes:= Quote{
		Title:  "Star wars trek",
		Body:   "Whenever you see a Fox, don't run",
		Author: "Gabriel",
	}
	renderWithData(w, "posted_quote.html", r, quotes)
}