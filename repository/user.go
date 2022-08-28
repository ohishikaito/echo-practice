package repository

import (
	database "github.com/ohishikaito/echo-practice/adapter/db"
	"github.com/ohishikaito/echo-practice/domain"
)

type (
	UserRepo interface {
		GetUsers() ([]*domain.User, error)
	}
	userRepo struct {
		db database.DB
	}
)

func NewUserRepo(db database.DB) UserRepo {
	return &userRepo{db}
}

func (r *userRepo) GetUsers() ([]*domain.User, error) {
	sql := `
SELECT *
FROM users
`
	rows, err := r.db.Queryx(sql)
	if err != nil {
		return nil, err
	}
	items := []*domain.User{}
	for rows.Next() {
		item := new(domain.User)
		if err := rows.StructScan(item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
