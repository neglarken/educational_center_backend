package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/neglarken/educational_center_backend/internal/app/store"
	"github.com/sirupsen/logrus"
)

type server struct {
	router mux.Router
	logger logrus.Logger
	store  store.Store
}

func NewServer(store store.Store) *server {
	s := &server{
		router: *mux.NewRouter(),
		logger: *logrus.New(),
		store:  store,
	}
	s.logger.Info("server is running on port 8080")
	s.configureUsersRouter()

	return s
}

func (s *server) NewSubRouter(str string) *mux.Router {
	return s.router.PathPrefix(str).Subrouter()
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) error(w http.ResponseWriter, r *http.Request, httpCode int, err error) {
	s.respond(w, r, httpCode, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, httpCode int, data interface{}) {
	w.WriteHeader(httpCode)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
