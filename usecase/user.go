package usecase

import (
	"github.com/ohishikaito/echo-practice/adapter/env"
	"github.com/ohishikaito/echo-practice/domain"
	"github.com/ohishikaito/echo-practice/repository"
)

type (
	UserUc interface {
		GetUsers() ([]*domain.User, error)
	}
	userUc struct {
		r repository.UserRepo
		e env.Env
	}
)

func NewUserUc(r repository.UserRepo, e env.Env) UserUc {
	return &userUc{r, e}
}

func (u *userUc) GetUsers() ([]*domain.User, error) {
	return u.r.GetUsers()
}
