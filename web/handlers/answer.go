package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jdanielnd/go-quiz/core/answer"
)

func MakeAnswerHandlers(r *mux.Router, n *negroni.Negroni, service answer.UseCase) {
	r.Handle("/v1/answers", n.With(
		negroni.Wrap(getAll(service)),
	)).Methods("GET", "OPTIONS")
	r.Handle("/v1/answers", n.With(
		negroni.Wrap(storeAnswer(service)),
	)).Methods("POST", "OPTIONS")
}

func getAll(service answer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		all, err := service.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(all)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError("Error while converting JSON"))
			return
		}
	})
}

func storeAnswer(service answer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var a answer.Answer
		err := json.NewDecoder(r.Body).Decode(&a)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(formatJSONError(err.Error()))
			return
		}

		err = service.Store(&a)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}
