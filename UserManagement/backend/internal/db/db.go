package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "./internal/db/users.db")
	if err != nil {
		return err
	}

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE,
			password TEXT
		);
	`)

	if err != nil {
		return err
	}

	fmt.Println("DB is initialized")

	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
