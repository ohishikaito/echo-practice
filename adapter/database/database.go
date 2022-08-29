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
	dBConnectionBuilder struct {
		dbUser      string
		dbPassword  string
		dbHost      string
		dbName      string
		dbOptions   string
		databaseUrl string
	}
)

func NewDBConnectionBuilder(env environment.Env) *dBConnectionBuilder {
	dbUser := env.GetDBUser()
	dbPassword := env.GetDBPassword()
	dbHost := env.GetDBHost()
	dbName := env.GetDBName()
	dbOptions := env.GetDBOptions()
	databaseUrl := dbUser + ":" + dbPassword + "@" + dbHost + "/" + dbName + dbOptions
	return &dBConnectionBuilder{
		dbUser,
		dbPassword,
		dbHost,
		dbName,
		dbOptions,
		databaseUrl,
	}
}

func NewDB(dBConnectionBuilder *dBConnectionBuilder) DB {
	db, err := sqlx.Connect("mysql", dBConnectionBuilder.databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	return DB{db}
}
