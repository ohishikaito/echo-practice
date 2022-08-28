package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/ohishikaito/echo-practice/adapter/env"
)

type (
	DB struct {
		*sqlx.DB
	}
)

func NewDB(e env.Env) DB {
	dbUser := e.GetDBUser()
	dbPassword := e.GetDBPassword()
	dbHost := e.GetDBHost()
	dbName := e.GetDBName()
	dbOptions := e.GetDBOptions()
	databaseUrl := dbUser + ":" + dbPassword + "@" + dbHost + "/" + dbName + dbOptions

	db, err := sqlx.Connect("mysql", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	return DB{db}
}
