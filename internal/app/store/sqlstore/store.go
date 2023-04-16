package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/neglarken/educational_center_backend/internal/app/store"
)

type Store struct {
	db              *sql.DB
	usersRepository *UsersRepository
	newsRepository  *NewsRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Users() store.UsersRepository {
	if s.usersRepository != nil {
		return s.usersRepository
	}
	s.usersRepository = &UsersRepository{
		store: s,
	}
	return s.usersRepository
}

func (s *Store) News() store.NewsRepository {
	if s.newsRepository != nil {
		return s.newsRepository
	}
	s.newsRepository = &NewsRepository{
		store: s,
	}
	return s.newsRepository
}
