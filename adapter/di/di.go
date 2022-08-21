package di

import (
	"github.com/ohishikaito/echo-practice/dject"
	"github.com/ohishikaito/echo-practice/repository"
	"github.com/ohishikaito/echo-practice/usecase"
)

func CreateContainer() (dject.Container, error) {
	container := dject.NewContainer()

	if err := container.Register(repository.NewUserRepo); err != nil {
		return nil, err
	}
	if err := container.Register(usecase.NewUserUc); err != nil {
		return nil, err
	}

	return container, nil
}
