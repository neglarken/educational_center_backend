package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/neglarken/educational_center_backend/internal/app/store"
	"github.com/sirupsen/logrus"
)

type ctxKey int8

const (
	ctxKeyRequestId ctxKey = iota
	ctxKeyUser      ctxKey = iota
)

type server struct {
	router       *mux.Router
	logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store
}

func NewServer(store store.Store, sessionStore sessions.Store) *server {
	s := &server{
		router:       mux.NewRouter(),
		logger:       logrus.New(),
		store:        store,
		sessionStore: sessionStore,
	}
	s.logger.Info("server is running on port 8080")
	s.configureRouters()

	return s
}

func (s *server) configureRouters() {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.configureUsersRouter() // /users
	s.configureNewsRouter()  // /news
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
