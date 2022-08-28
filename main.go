package main

import (
	"log"

	"github.com/ohishikaito/echo-practice/adapter/db"
	"github.com/ohishikaito/echo-practice/adapter/di"
	"github.com/ohishikaito/echo-practice/adapter/env"
	"github.com/ohishikaito/echo-practice/adapter/server"
)

func main() {
	env := env.NewEnv()
	db := db.NewDB(env)
	container, err := di.CreateContainer(env, db)
	if err != nil {
		log.Fatal(err)
	}
	r := server.NewServer(env.GetPort(), container)
	r.Serve()
}
