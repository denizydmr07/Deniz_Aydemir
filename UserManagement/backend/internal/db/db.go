package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// DB is the database connection
var DB *sql.DB

// InitDB initializes the database
// It returns an error if something goes wrong
// It prints a message to the console if the database is initialized successfully
func InitDB() error {

	// Open the database
	var err error
	DB, err = sql.Open("sqlite3", "./internal/db/users.db")

	// Check Open Error
	if err != nil {
		return err
	}

	// Create users table if it does not exist
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE,
			password TEXT
		);
	`)

	// Check Exec Error
	if err != nil {
		return err
	}

	// Print message to the console
	fmt.Println("DB is initialized")

	return nil
}

// CloseDB closes the database connection
// It prints a message to the console if the database is closed successfully
func CloseDB() {
	if DB != nil {
		DB.Close()
		fmt.Println("DB is closed")
	}
}
