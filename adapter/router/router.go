package router

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/ohishikaito/echo-practice/dject"
)

type (
	router struct {
		e *echo.Echo
	}
	Router interface {
		Serve()
	}
)

var Container dject.Container

func NewRouter(container dject.Container) Router {
	Container = container
	e := echo.New()
	e.GET("/users", getUsers)
	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)
	return &router{e}
}

func (r *router) Serve() {
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if err := r.e.Start(addr); err != nil {
		log.Fatal(err)
	}
}
