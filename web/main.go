package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jdanielnd/go-quiz/core/question"
	"github.com/jdanielnd/go-quiz/web/handlers"
)

func main() {
	db, err := sql.Open("sqlite3", "data/quiz.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	service := question.NewService(db)
	r := mux.NewRouter().StrictSlash(true)
	n := negroni.New(
		negroni.NewLogger(),
	)

	handlers.MakeQuestionHandlers(r, n, service)

	http.Handle("/", r)

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":4000",
		Handler:      http.DefaultServeMux,
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
