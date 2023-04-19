package apiserver

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/neglarken/educational_center_backend/internal/app/model"
)

func (s *server) configureNewsRouter() {
	sub := s.NewSubRouter("/news")
	sub.Use(s.authUser)
	sub.HandleFunc("/", s.HandleGetNews()).Methods("GET")
	sub.HandleFunc("/count", s.HandleGetNewsUnreadCount()).Methods(("GET"))
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
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		u := r.Context().Value(ctxKeyUser).(*model.Users)
		_, err = s.store.NewsUsers().FindByIds(intId, u.Id)
		if err != nil && err != sql.ErrNoRows {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		if err == sql.ErrNoRows {
			if err := s.store.NewsUsers().ReadNewsById(intId, u.Id); err != nil {
				s.error(w, r, http.StatusInternalServerError, err)
				return
			}
		}
		s.respond(w, r, http.StatusOK, n)
	}
}

func (s *server) HandleGetNewsReadCount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := r.Context().Value(ctxKeyUser).(*model.Users)
		c, err := s.store.NewsUsers().GetCountById(u.Id)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, c)
	}
}

func (s *server) HandleGetNewsUnreadCount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := r.Context().Value(ctxKeyUser).(*model.Users)
		countRead, err := s.store.NewsUsers().GetCountById(u.Id)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		countAll, err := s.store.News().GetCount()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		countResult := &model.CountOfNews{Count: countAll.Count - countRead.Count}
		s.respond(w, r, http.StatusOK, countResult)
	}
}
