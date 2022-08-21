package di

import (
	"reflect"

	"github.com/ohishikaito/echo-practice/adapter/env"
	"github.com/ohishikaito/echo-practice/dject"
	"github.com/ohishikaito/echo-practice/repository"
	"github.com/ohishikaito/echo-practice/usecase"
)

func CreateContainer(e env.Env) (dject.Container, error) {
	container := dject.NewContainer()
	// env opt
	envOpt := dject.RegisterOptions{Interfaces: []reflect.Type{reflect.TypeOf((*env.Env)(nil)).Elem()}}
	if err := container.Register(e, envOpt); err != nil {
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
