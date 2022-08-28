package usecase

import (
	"github.com/ohishikaito/echo-practice/adapter/environment"
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
		env    environment.Env
		paypal paypal.PaypalClient
	}
)

func NewUserUc(r repository.UserRepo, env environment.Env, paypal paypal.PaypalClient) UserUc {
	return &userUc{r, env, paypal}
}

func (u *userUc) GetUsers() ([]*domain.User, error) {
	return u.r.GetUsers()
}
