package quooote

import (
	"net/http"
)

func Index (w http.ResponseWriter, r *http.Request) {
	render(w, "index.html", r)
}
