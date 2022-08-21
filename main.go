package main

import (
	"log"

	"github.com/ohishikaito/echo-practice/adapter/di"
	"github.com/ohishikaito/echo-practice/adapter/server"
)

func main() {
	container, err := di.CreateContainer()
	if err != nil {
		log.Fatal(err)
	}
	r := server.NewServer(container)
	r.Serve()
}
