package quooote

import (
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

func render (w http.ResponseWriter, tmpl string,r *http.Request) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		logrus.Fatalln(err)
	}
	_ = t.Execute(w, nil)
}
