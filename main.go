package main

import (
	"log"
	"os"

	"github.com/ohishikaito/echo-practice/adapter/di"
	"github.com/ohishikaito/echo-practice/adapter/env"
	"github.com/ohishikaito/echo-practice/adapter/server"
)

func main() {
	env := env.NewEnv()
	container, err := di.CreateContainer(env)
	if err != nil {
		log.Fatal(err)
	}
	r := server.NewServer(os.Getenv("PORT"), container)
	r.Serve()
}
