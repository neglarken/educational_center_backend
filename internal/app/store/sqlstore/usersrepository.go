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

func (r *UsersRepository) Edit(u *model.Users, id int) error {
	return r.store.db.QueryRow(
		"UPDATE users SET first_name = $1, last_name = $2, surname = $3, phone_number = $4 WHERE id = $5 RETURNING id",
		u.FirstName,
		u.LastName,
		u.Surname,
		u.PhoneNumber,
		id,
	).Err()
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

func (r *UsersRepository) FindById(id int) (*model.Users, error) {
	u := &model.Users{}
	if err := r.store.db.QueryRow(
		"SELECT id, login, password, first_name, last_name, surname, phone_number FROM users WHERE id = $1",
		id,
	).Scan(
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
