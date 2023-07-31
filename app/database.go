package app

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "my_database.db")
	if err != nil {
		log.Fatal(err)
	}
	createTable()
}

func createTable() {
	sqlStmt := `
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT,
			body TEXT,
			userId INTEGER
		);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return db
}
