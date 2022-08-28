package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/ohishikaito/echo-practice/adapter/environment"
)

type (
	DB struct {
		*sqlx.DB
	}
)

func NewDB(env environment.Env) DB {
	dbUser := env.GetDBUser()
	dbPassword := env.GetDBPassword()
	dbHost := env.GetDBHost()
	dbName := env.GetDBName()
	dbOptions := env.GetDBOptions()
	databaseUrl := dbUser + ":" + dbPassword + "@" + dbHost + "/" + dbName + dbOptions

	db, err := sqlx.Connect("mysql", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	return DB{db}
}
