package environment

import (
	"os"
)

type (
	Env interface {
		GetPort() string
		GetDBUser() string
		GetDBPassword() string
		GetDBHost() string
		GetDBName() string
		GetDBOptions() string
	}
	env struct {
		port       string
		dbUser     string
		dbPassword string
		dbHost     string
		dbName     string
		dbOptions  string
	}
)

func NewEnv() Env {
	port := os.Getenv("PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbOptions := os.Getenv("DB_OPTIONS")
	return &env{
		port,
		dbUser,
		dbPassword,
		dbHost,
		dbName,
		dbOptions,
	}
}

func (e *env) GetPort() string {
	return e.port
}
func (e *env) GetDBUser() string {
	return e.dbUser
}
func (e *env) GetDBPassword() string {
	return e.dbPassword
}
func (e *env) GetDBHost() string {
	return e.dbHost
}
func (e *env) GetDBName() string {
	return e.dbName
}
func (e *env) GetDBOptions() string {
	return e.dbOptions
}
