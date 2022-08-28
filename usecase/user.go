package usecase

import (
	"github.com/ohishikaito/echo-practice/adapter/env"
	"github.com/ohishikaito/echo-practice/adapter/paypal"
	"github.com/ohishikaito/echo-practice/domain"
	"github.com/ohishikaito/echo-practice/repository"
)

type (
	UserUc interface {
		GetUsers() ([]*domain.User, error)
	}
	userUc struct {
		r      repository.UserRepo
		e      env.Env
		paypal paypal.PaypalClient
	}
)

// func NewUserUc(r repository.UserRepo, e env.Env) UserUc {
// 	return &userUc{r, e}
// }

func NewUserUc(r repository.UserRepo, e env.Env, paypal paypal.PaypalClient) UserUc {
	return &userUc{r, e, paypal}
}

func (u *userUc) GetUsers() ([]*domain.User, error) {
	return u.r.GetUsers()
}
