package store

import "github.com/neglarken/educational_center_backend/internal/app/model"

type UsersRepository interface {
	Create(*model.Users) error
	FindByLogin(string) (*model.Users, error)
	Find(int) (*model.Users, error)
}
