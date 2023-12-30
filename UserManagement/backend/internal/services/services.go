package services

import (
	"encoding/json"
	"fmt"

	"github.com/denizydmr07/UserManagement/backend/internal/db"
)

func InsertUser(username, password string) error {
	_, err := db.DB.Exec(`
		INSERT INTO users (username, password)
		VALUES (?, ?);
	`, username, password)
	// if username
	if err != nil {
		return err
	}

	return nil
}

// GetUsers returns all users in a JSON format
func GetUsers() ([]byte, error) {
	rows, err := db.DB.Query(`
		SELECT * FROM users;
	`)
	if err != nil {
		return nil, err
	}

	var users []map[string]interface{}

	for rows.Next() {
		var id int
		var username, password string

		err := rows.Scan(&id, &username, &password)
		if err != nil {
			return nil, err
		}

		user := map[string]interface{}{
			"id":       id,
			"username": username,
			"password": password,
		}

		users = append(users, user)
	}

	usersJSON, err := json.Marshal(users)

	if err != nil {
		return nil, err
	}

	return usersJSON, nil
}

func GetUserWithID(id string) ([]byte, error) {
	row := db.DB.QueryRow(`
		SELECT * FROM users WHERE id = ?;
	`, id)

	var username, password string

	err := row.Scan(&id, &username, &password)
	if err != nil {
		return nil, err
	}

	user := map[string]interface{}{
		"id":       id,
		"username": username,
		"password": password,
	}

	userJSON, err := json.Marshal(user)

	if err != nil {
		return nil, err
	}

	return userJSON, nil
}

func UpdateUser(id, username, password string) error {
	_, err := db.DB.Exec(`
		UPDATE users
		SET username = ?, password = ?
		WHERE id = ?;
	`, username, password, id)

	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(id string) error {
	fmt.Println(id)
	_, err := db.DB.Exec(`
		DELETE FROM users WHERE id = ?;
	`, id)

	if err != nil {
		return err
	}

	return nil
}
