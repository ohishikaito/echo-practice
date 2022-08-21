package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type (
	DB struct {
		*sqlx.DB
	}
)

func NewDB() DB {
	db, err := sqlx.Connect("mysql", "user=foo dbname=bar")
	if err != nil {
		log.Fatal(err)
	}
	return DB{db}
}
