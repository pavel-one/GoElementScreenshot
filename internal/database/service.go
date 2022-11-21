package database

import (
	_ "embed"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

//go:embed schema/schema.sql
var schema string

func GetDatabase() (db *sqlx.DB, err error) {
	db, err = sqlx.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)

	return db, err
}
