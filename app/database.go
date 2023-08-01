package app

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/dhaaana/go-http-server/config"
	"github.com/dhaaana/go-http-server/utils"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	// initialize the database (sqlite)
	// var err error
	// db, err = sql.Open("sqlite3", "my_database.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// initialize the database (postgresql)
	dbVariables := map[string]string{
		"DB_HOST":     "",
		"DB_PORT":     "",
		"DB_USER":     "",
		"DB_PASSWORD": "",
		"DB_NAME":     "",
	}

	for key := range dbVariables {
		value, err := config.GetEnvVariables(key)
		if err != nil {
			utils.LogError("Error getting database variables:", err)
			os.Exit(1)
		}
		dbVariables[key] = value
	}

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbVariables["DB_HOST"],
		dbVariables["DB_PORT"],
		dbVariables["DB_USER"],
		dbVariables["DB_PASSWORD"],
		dbVariables["DB_NAME"],
	)

	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		utils.LogError("Error connecting to database:", err)
		os.Exit(1)
	}

	createTable()
	runSeeder()
}

func createTable() {
	sqlStmt := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			password TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL,
			name TEXT NOT NULL
		);
		
		CREATE TABLE IF NOT EXISTS posts (
			id SERIAL PRIMARY KEY,
			title TEXT,
			body TEXT,
			userId INTEGER REFERENCES users(id) ON DELETE CASCADE
		);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		utils.LogError("Error creating table:", err)
		os.Exit(1)
	}
}

func runSeeder() {
	if shouldSeed, _ := config.GetEnvBool("SEED_DB"); shouldSeed {
		err := SeedData()
		if err != nil {
			utils.LogError("Error seeding database:", err)
		}
	}
}

func GetDB() *sql.DB {
	return db
}
