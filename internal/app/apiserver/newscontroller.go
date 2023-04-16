package apiserver

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *server) configureNewsRouter() {
	sub := s.NewSubRouter("/news")
	sub.HandleFunc("/", s.HandleGetNews()).Methods("GET")
	sub.HandleFunc("/{id}", s.HandleGetNewsById()).Methods(("GET"))
}

func (s *server) HandleGetNews() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		n, err := s.store.News().Get()
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusOK, n)
	}
}

func (s *server) HandleGetNewsById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			s.error(w, r, http.StatusBadGateway, errors.New("id is missing in parameters"))
			return
		}
		intId, err := strconv.Atoi(id)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		n, err := s.store.News().GetById(intId)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusOK, n)
	}
}
