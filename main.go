package main

import (
	"github.com/GabielFemi/quooote/quooote"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", quooote.Index)
	server := &http.Server{
		Addr: "127.0.0.1:8000",
		Handler: router,
	}

	logrus.Fatalln(server.ListenAndServe())
}