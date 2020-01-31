package model

import (
	"database/sql"
	"log"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/todo")
	if err != nil {
		log.Fatal(err)
	}
	con = db
	return db
}
