package store

import (
	"github.com/neglarken/educational_center_backend/internal/app/model"
)

type UsersRepository interface {
	Create(*model.Users) error
	Edit(*model.Users, int) error
	FindByLogin(string) (*model.Users, error)
	FindById(int) (*model.Users, error)
}

type NewsRepository interface {
	Get() ([]*model.News, error)
	GetById(int) (*model.News, error)
	GetCount() (*model.CountOfNews, error)
}

type NewsUsersRepository interface {
	GetCountById(int) (*model.CountOfNews, error)
	ReadNewsById(int, int) error
	FindByIds(int, int) (*model.NewsUsers, error)
}
