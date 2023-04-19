package sqlstore

import "github.com/neglarken/educational_center_backend/internal/app/model"

type NewsRepository struct {
	store *Store
}

func (r *NewsRepository) Get() ([]*model.News, error) {
	ns := make([]*model.News, 0)
	rows, err := r.store.db.Query("SELECT * FROM news")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		n := &model.News{}
		rows.Scan(
			&n.Id,
			&n.Title,
			&n.Description,
			&n.CreatedAt,
		)
		ns = append(ns, n)
	}
	return ns, nil
}

func (r *NewsRepository) GetById(id int) (*model.News, error) {
	n := &model.News{}
	if err := r.store.db.QueryRow(
		"SELECT * FROM news WHERE id = $1",
		id,
	).Scan(
		&n.Id,
		&n.Title,
		&n.Description,
		&n.CreatedAt,
	); err != nil {
		return nil, err
	}

	return n, nil
}

func (r *NewsRepository) GetCount() (*model.CountOfNews, error) {
	count := &model.CountOfNews{}
	if err := r.store.db.QueryRow(
		"SELECT count(*) FROM news",
	).Scan(
		&count.Count,
	); err != nil {
		return nil, err
	}
	return count, nil
}
