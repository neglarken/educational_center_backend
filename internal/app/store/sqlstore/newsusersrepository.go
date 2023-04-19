package sqlstore

import "github.com/neglarken/educational_center_backend/internal/app/model"

type NewsUsersRepository struct {
	store *Store
}

func (r *NewsUsersRepository) GetCountById(id int) (*model.CountOfNews, error) {
	c := &model.CountOfNews{}
	if err := r.store.db.QueryRow(
		"SELECT count(*) FROM news_users WHERE user_id = $1",
		id,
	).Scan(
		&c.Count,
	); err != nil {
		return nil, err
	}
	return c, nil
}

func (r *NewsUsersRepository) FindByIds(NewsId int, UserId int) (*model.NewsUsers, error) {
	nu := &model.NewsUsers{}
	if err := r.store.db.QueryRow("SELECT * FROM news_users WHERE user_id = $1 AND news_id = $2", UserId, NewsId).Scan(
		&nu.UserId,
		&nu.NewsId,
	); err != nil {
		return nil, err
	}
	return nu, nil
}

func (r *NewsUsersRepository) ReadNewsById(NewsId int, UserId int) error {
	n := &model.NewsUsers{
		NewsId: NewsId,
		UserId: UserId,
	}
	return r.store.db.QueryRow(
		"INSERT INTO news_users (user_id, news_id) VALUES ($1, $2) RETURNING news_id, user_id",
		n.UserId,
		n.NewsId,
	).Scan(&n.NewsId, &n.UserId)
}
