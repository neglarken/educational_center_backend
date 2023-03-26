package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/neglarken/educational_center_backend/internal/app/store/sqlstore"
)

func Start(config *Config) error {
	db, err := NewDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	s := NewServer(store)
	return http.ListenAndServe(config.BindAddr, s)
}

func NewDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
