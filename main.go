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
	router.PathPrefix("/assets").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	router.HandleFunc("/", quooote.Index)
	router.HandleFunc("/quotes", quooote.FetchQuotes).Methods("GET")
	router.HandleFunc("/quotes/add", quooote.AddQuote).Methods("GET")
	router.HandleFunc("/quotes/add", quooote.PostQuote).Methods("POST")



	server := &http.Server{
		Addr: "127.0.0.1:8000",
		Handler: router,
	}
	logrus.Infoln("App is listening on ", server.Addr)
	logrus.Fatalln(server.ListenAndServe())
}
