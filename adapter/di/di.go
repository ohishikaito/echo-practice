package di

import (
	"reflect"

	"github.com/ohishikaito/echo-practice/adapter/database"
	"github.com/ohishikaito/echo-practice/adapter/environment"
	"github.com/ohishikaito/echo-practice/adapter/paypal"
	"github.com/ohishikaito/echo-practice/dject"
	"github.com/ohishikaito/echo-practice/repository"
	"github.com/ohishikaito/echo-practice/usecase"
)

func CreateContainer(env environment.Env, db database.DB) (dject.Container, error) {
	container := dject.NewContainer()

	envOpt := dject.RegisterOptions{Interfaces: []reflect.Type{reflect.TypeOf((*environment.Env)(nil)).Elem()}}
	if err := container.Register(env, envOpt); err != nil {
		return nil, err
	}
	paypalOpt := dject.RegisterOptions{Interfaces: []reflect.Type{reflect.TypeOf((*paypal.PaypalClient)(nil)).Elem()}}
	if err := container.Register(paypal.NewPaypalClient(env), paypalOpt); err != nil {
		return nil, err
	}
	dbOpt := dject.RegisterOptions{Interfaces: []reflect.Type{reflect.TypeOf((*database.DB)(nil)).Elem()}}
	if err := container.Register(db, dbOpt); err != nil {
		return nil, err
	}

	if err := container.Register(repository.NewUserRepo); err != nil {
		return nil, err
	}
	if err := container.Register(usecase.NewUserUc); err != nil {
		return nil, err
	}

	return container, nil
}
