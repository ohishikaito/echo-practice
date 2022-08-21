package di

import (
	"github.com/ohishikaito/echo-practice/dject"
	"github.com/ohishikaito/echo-practice/repository"
	"github.com/ohishikaito/echo-practice/usecase"
)

var Container dject.Container

func CreateContainer() (dject.Container, error) {
	container := dject.NewContainer()
	Container = container

	// user queries
	if err := container.Register(repository.NewUserRepo); err != nil {
		return nil, err
	}
	if err := container.Register(usecase.NewUserUc); err != nil {
		return nil, err
	}

	return container, nil
}
