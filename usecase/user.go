package usecase

import (
	"github.com/ohishikaito/echo-practice/domain"
	"github.com/ohishikaito/echo-practice/repository"
)

type (
	UserUc interface {
		GetUsers() ([]*domain.User, error)
	}
	userUc struct {
		r repository.UserRepo
	}
)

func NewUserUc(r repository.UserRepo) UserUc {
	return &userUc{r}
}

func (u *userUc) GetUsers() ([]*domain.User, error) {
	return u.r.GetUsers()
}
