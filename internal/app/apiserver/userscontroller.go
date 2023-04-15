package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/neglarken/educational_center_backend/internal/app/model"
)

func (s *server) configureUsersRouter() {
	sub := s.NewSubRouter("/users")
	sub.HandleFunc("/register", s.HandleUsersCreate()).Methods("POST")
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
