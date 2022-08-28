package main

import (
	"log"

	"github.com/ohishikaito/echo-practice/adapter/database"
	"github.com/ohishikaito/echo-practice/adapter/di"
	"github.com/ohishikaito/echo-practice/adapter/environment"
	"github.com/ohishikaito/echo-practice/adapter/server"
)

func main() {
	env := environment.NewEnv()
	db := database.NewDB(env)
	container, err := di.CreateContainer(env, db)
	if err != nil {
		log.Fatal(err)
	}
	r := server.NewServer(env.GetPort(), container)
	r.Serve()
}
