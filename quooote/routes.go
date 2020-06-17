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