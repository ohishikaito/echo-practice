package env

import (
	"os"
)

type (
	Env interface {
		GetPort() string
	}
	env struct {
		port string
	}
)

func NewEnv() Env {
	port := os.Getenv("PORT")
	return &env{
		port,
	}
}

func (e *env) GetPort() string {
	return e.port
}
