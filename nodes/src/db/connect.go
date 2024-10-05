package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// i just need to connect with the db here

func Connect() (database *sql.DB) {
	database, _ = sql.Open("sqlite3", "./db/database.db")
	return
}
