package quooote

import (
	"net/http"
)

func Index (w http.ResponseWriter, r *http.Request) {
	render(w, "index.html", r)
}

func AddQuote (w http.ResponseWriter, r *http.Request) {
	render(w, "add_quote.html", r)
}

func PostQuote (w http.ResponseWriter, r *http.Request) {
	render(w, "posted_quote.html", r)
}