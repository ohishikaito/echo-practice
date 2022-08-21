package server

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/ohishikaito/echo-practice/adapter/serverCtx"
	"github.com/ohishikaito/echo-practice/controller"
	"github.com/ohishikaito/echo-practice/dject"
)

type (
	server struct {
		e    *echo.Echo
		port string
	}
	Server interface {
		Serve()
	}
)

func NewServer(port string, container dject.Container) Server {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &serverCtx.ServerCtx{
				Context:   c,
				Container: container,
			}
			return next(cc)
		}
	})
	e.GET("/users", controller.GetUsers)
	// e.POST("/users", saveUser)
	// e.GET("/users/:id", getUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)
	return &server{e, port}
}

func (r *server) Serve() {
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if err := r.e.Start(addr); err != nil {
		log.Fatal(err)
	}
}
