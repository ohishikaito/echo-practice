package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/ohishikaito/echo-practice/adapter/di"
	"github.com/ohishikaito/echo-practice/domain"
	"github.com/ohishikaito/echo-practice/usecase"
)

func main() {
	e := echo.New()
	e.GET("/users", getUsers)
	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)
	e.GET("/hoge", hoge)

	address := fmt.Sprintf(":%s", os.Getenv("PORT"))
	e.Logger.Fatal(e.Start(address))
}

func hoge(c echo.Context) error {
	id := c.Param("hoge")
	fmt.Println(id)
	return nil
}

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func getUsers(c echo.Context) error {
	// r := repository.NewUserRepo()
	// uc := usecase.NewUserUc(r)
	// users, err := uc.GetUsers()
	// if err != nil {
	// 	return err
	// }
	// return c.JSON(200, users)

	container, err := di.CreateContainer()
	res := []*domain.User{}
	if err != nil {
		log.Fatal(err)
	}
	if err := container.Invoke(func(
		usecase usecase.UserUc,
	) error {
		rslt, err := usecase.GetUsers()
		if err != nil {
			return err
		}
		res = rslt
		return nil
	}); err != nil {
		return err
	}
	return c.JSON(200, res)
}

func saveUser(c echo.Context) error {
	// Get name
	name := c.FormValue("name")
	// Get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}
	fmt.Println(avatar.Filename)

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}
