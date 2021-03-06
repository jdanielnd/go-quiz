package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jdanielnd/go-quiz/core/alternative"
	"github.com/jdanielnd/go-quiz/core/answer"
	"github.com/jdanielnd/go-quiz/core/question"
	"github.com/jdanielnd/go-quiz/web/handlers"
)

func main() {
	db, err := sql.Open("sqlite3", "data/quiz.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	questionService := question.NewService(db)
	alternativeService := alternative.NewService(db)
	answerService := answer.NewService(db)

	r := mux.NewRouter().StrictSlash(true)
	n := negroni.New(
		negroni.NewLogger(),
	)

	handlers.MakeQuestionHandlers(r, n, questionService)
	handlers.MakeAlternativeHandlers(r, n, alternativeService)
	handlers.MakeAnswerHandlers(r, n, answerService)

	http.Handle("/", r)

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":" + port,
		Handler:      http.DefaultServeMux,
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
