package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/neglarken/educational_center_backend/internal/app/model"
)

const (
	sessionName = "educational_center"
)

func (s *server) configureUsersRouter() {
	sub := s.NewSubRouter("/users")
	sub.HandleFunc("/register", s.HandleUsersCreate()).Methods("POST")
	sub.HandleFunc("/sessions", s.HandleSessionsCreate()).Methods("POST")

	// /private/...
	private := s.router.PathPrefix("/private").Subrouter()
	private.Use(s.authUser)
	private.HandleFunc("/whoami", s.handleWhoAmI()).Methods("GET")
}

func (s *server) HandleUsersCreate() http.HandlerFunc {
	type request struct {
		Login               string "json:\"login\""
		UnencryptedPassword string "json:\"password\""
		FirstName           string "json:\"first_name\""
		LastName            string "json:\"last_name\""
		Surname             string "json:\"surname\""
		PhoneNumber         string "json:\"phone_number\""
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &model.Users{
			Login:               req.Login,
			UnencryptedPassword: req.UnencryptedPassword,
			FirstName:           req.FirstName,
			LastName:            req.LastName,
			Surname:             req.Surname,
			PhoneNumber:         req.PhoneNumber,
		}
		if err := s.store.Users().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) HandleSessionsCreate() http.HandlerFunc {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u, err := s.store.Users().FindByLogin(req.Login)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errors.New("Invalid login or password"))
			return
		}
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		session.Values["user_id"] = u.Id
		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) handleWhoAmI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*model.Users))
	}
}
