package main

import (
	"log"

	"github.com/ohishikaito/echo-practice/adapter/di"
	"github.com/ohishikaito/echo-practice/adapter/router"
)

func main() {
	container, err := di.CreateContainer()
	if err != nil {
		log.Fatal(err)
	}
	r := router.NewRouter(container)
	r.Serve()
}
