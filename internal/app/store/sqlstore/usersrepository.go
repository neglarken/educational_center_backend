package sqlstore

import "github.com/neglarken/educational_center_backend/internal/app/model"

type UsersRepository struct {
	store *Store
}

func (r *UsersRepository) Create(u *model.Users) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (login, password, first_name, last_name, surname, phone_number) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		u.Login,
		u.Password,
		u.FirstName,
		u.LastName,
		u.Surname,
		u.PhoneNumber,
	).Scan(&u.Id)
}

func (r *UsersRepository) FindByLogin(login string) (*model.Users, error) {
	u := &model.Users{}
	if err := r.store.db.QueryRow("SELECT * FROM users WHERE login = $1", login).Scan(
		&u.Id,
		&u.Login,
		&u.Password,
		&u.FirstName,
		&u.LastName,
		&u.Surname,
		&u.PhoneNumber,
	); err != nil {
		return nil, err
	}
	return u, nil
}
