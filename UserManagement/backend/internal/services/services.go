package services

import (
	"encoding/json"
	"fmt"

	"github.com/denizydmr07/UserManagement/backend/internal/db"
)

// InsertUser inserts a new user to the database
// It takes username and password as parameters
// It returns an error if something goes wrong
// It returns nil if the user is inserted successfully
func InsertUser(username, password string) error {

	// Insert user to the database
	_, err := db.DB.Exec(`
		INSERT INTO users (username, password)
		VALUES (?, ?);
	`, username, password)

	// Check Exec Error
	if err != nil {
		return err
	}

	return nil
}

// GetUsers returns all users in a JSON format
// It returns an error if something goes wrong
// It returns nil if users are retrieved successfully
func GetUsers() ([]byte, error) {

	// Get users from the database
	rows, err := db.DB.Query(`
		SELECT * FROM users;
	`)

	// Check Query Error
	if err != nil {
		return nil, err
	}

	// Create a slice of users
	var users []map[string]interface{}

	// Iterate over rows
	for rows.Next() {

		// Create variables to store row data
		var id int
		var username, password string

		// Scan rows and store row data in variables
		err := rows.Scan(&id, &username, &password)
		if err != nil {
			return nil, err
		}

		// Create a user map
		user := map[string]interface{}{
			"id":       id,
			"username": username,
			"password": password,
		}

		// Append user to users slice
		users = append(users, user)
	}

	// Convert users slice to JSON
	usersJSON, err := json.Marshal(users)

	// Check Marshal Error
	if err != nil {
		return nil, err
	}

	// Return usersJSON
	return usersJSON, nil
}

// GetUserWithID returns a user with the given id in a JSON format
// It takes id as a parameter
// It returns an error if something goes wrong
// It returns nil if the user is retrieved successfully
func GetUserWithID(id string) ([]byte, error) {

	// Get user from the database
	row := db.DB.QueryRow(`
		SELECT * FROM users WHERE id = ?;
	`, id)

	// Create variables to store row data
	var username, password string

	// Scan row and store row data in variables
	err := row.Scan(&id, &username, &password)

	// Check Scan Error
	if err != nil {
		return nil, err
	}

	// Create a user map
	user := map[string]interface{}{
		"id":       id,
		"username": username,
		"password": password,
	}

	// Convert user map to JSON
	userJSON, err := json.Marshal(user)

	// Check Marshal Error
	if err != nil {
		return nil, err
	}

	// Return userJSON
	return userJSON, nil
}

// UpdateUser updates a user with the given id in the database
// It takes id, username and password as parameters
// It returns an error if something goes wrong
func UpdateUser(id, username, password string) error {

	// Update user in the database
	_, err := db.DB.Exec(`
		UPDATE users
		SET username = ?, password = ?
		WHERE id = ?;
	`, username, password, id)

	// Check Exec Error
	if err != nil {
		return err
	}

	// Return nil
	return nil
}

// DeleteUser deletes a user with the given id from the database
// It takes id as a parameter
// It returns an error if something goes wrong
func DeleteUser(id string) error {

	// Delete user from the database
	fmt.Println(id)
	_, err := db.DB.Exec(`
		DELETE FROM users WHERE id = ?;
	`, id)

	// Check Exec Error
	if err != nil {
		return err
	}

	return nil
}
