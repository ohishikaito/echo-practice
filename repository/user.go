package repository

import "github.com/ohishikaito/echo-practice/domain"

type (
	UserRepo interface {
		GetUsers() ([]*domain.User, error)
	}
	userRepo struct{}
)

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (r *userRepo) GetUsers() ([]*domain.User, error) {
	users := []*domain.User{}
	users = append(users, &domain.User{
		Id:   2,
		Name: "name",
	})
	return users, nil
}
