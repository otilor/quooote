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

func renderWithData(w http.ResponseWriter, tmpl string, r *http.Request, data interface{}) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		logrus.Fatalln(err)
	}

	err = t.Execute(w, data)
	if err != nil {
		logrus.Fatalln(err)
	}
}