package database

import (
	"awesomeProject2/settings"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var Link *sql.DB

func Connect(options *settings.Setting) error {
	var e error
	Link, e = sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		options.DbHost,
		options.DbPort,
		options.DbUser,
		options.DbPass,
		options.DbName,
	))
	if e != nil {
		return e
	}

	e = Link.Ping()
	if e != nil {
		return e
	}

	return nil
}
