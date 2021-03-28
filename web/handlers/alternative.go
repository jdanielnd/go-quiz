package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jdanielnd/go-quiz/core/alternative"
)

func MakeAlternativeHandlers(r *mux.Router, n *negroni.Negroni, service alternative.UseCase) {
	r.Handle("/v1/questions/{id}/alternatives", n.With(
		negroni.Wrap(getAllAlternatives(service)),
	)).Methods("GET", "OPTIONS")
	// r.Handle("/v1/questions/{id}", n.With(
	// 	negroni.Wrap(getQuestion(service)),
	// )).Methods("GET", "OPTIONS")
	r.Handle("/v1/questions/{id}/alternatives", n.With(
		negroni.Wrap(storeAlternative(service)),
	)).Methods("POST", "OPTIONS")
	// r.Handle("/v1/questions/{id}", n.With(
	// 	negroni.Wrap(updateQuestion(service)),
	// )).Methods("PUT", "OPTIONS")
	// r.Handle("/v1/questions/{id}", n.With(
	// 	negroni.Wrap(removeQuestion(service)),
	// )).Methods("DELETE", "OPTIONS")
}

func getAllAlternatives(service alternative.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)

		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(formatJSONError(err.Error()))
			return
		}

		all, err := service.GetAllFromQuestion(id)
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

func storeAlternative(service alternative.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var a alternative.Alternative
		err := json.NewDecoder(r.Body).Decode(&a)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(formatJSONError(err.Error()))
			return
		}

		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(formatJSONError(err.Error()))
			return
		}
		if id != a.QuestionID {
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
