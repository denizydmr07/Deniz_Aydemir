package controllers

import (
	"fmt"
	"net/http"

	"github.com/denizydmr07/UserManagement/backend/internal/services"
)

// SaveUser saves a user to the database
// It takes username and password from the request body
// It returns a response with status code and message
// It prints the status code and message to the console
func SaveUser(w http.ResponseWriter, r *http.Request) {

	// Get username and password from request body
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Check username and password
	if username == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("username and password cannot be empty"))
		return
	}

	// Insert user to the database
	err := services.InsertUser(username, password)

	// Check InsertUser Error
	if err != nil {

		// Check if username already exists
		if err.Error() == "UNIQUE constraint failed: users.username" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("username already exists"))
			fmt.Println("username already exists")
		} else {

			// Internal Server Error
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
			fmt.Println("internal server error")
		}
		return
	}

	// User saved successfully
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user saved successfully"))
	fmt.Println("user saved successfully")
}

// GetUsers retrieves all users from the database
// It returns a response with status code and message
// It prints the status code and message to the console
func GetUsers(w http.ResponseWriter, r *http.Request) {

	// Get users from the database
	users, err := services.GetUsers()
	if err != nil {

		// Internal Server Error
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		fmt.Println("internal server error")
		return
	}

	// Users retrieved successfully
	w.WriteHeader(http.StatusOK)
	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")
	w.Write(users)
	fmt.Println("users retrieved successfully")
}

// GetUserWithID retrieves a user from the database with the given id
// It takes id from the request body
// It returns a response with status code and message
// It prints the status code and message to the console
func GetUserWithID(w http.ResponseWriter, r *http.Request) {

	// Get id from request body
	id := r.FormValue("id")

	// Check id
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("id cannot be empty"))
		return
	}

	// Get user from the database
	user, err := services.GetUserWithID(id)

	// Check GetUserWithID Error
	if err != nil {

		// Check if user not found
		if err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("user not found"))
			fmt.Println("user not found")
		} else {
			// Internal Server Error
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
			fmt.Println("internal server error")
		}
		return
	}

	// User retrieved successfully
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(user)
	fmt.Println("user retrieved successfully")
}

// UpdateUser updates a user in the database with the given id
// It takes id, username and password from the request body
// It returns a response with status code and message
// It prints the status code and message to the console
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	// Get id, username and password from request body
	id := r.FormValue("id")
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Check id, username and password
	if id == "" || username == "" || password == "" {

		// Bad Request
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("id, username and password cannot be empty"))
		return
	}

	// Update user in the database
	err := services.UpdateUser(id, username, password)

	// Check UpdateUser Error
	if err != nil {

		// Check if user not found
		if err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("user not found"))
			fmt.Println("user not found")
		} else {

			// Internal Server Error
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
			fmt.Println("internal server error")
		}
		return
	}

	// User updated successfully
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user updated successfully"))
	fmt.Println("user updated successfully")
}

// DeleteUser deletes a user in the database with the given id
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	// Get id from request body
	id := r.FormValue("id")

	// Check id
	if id == "" {

		// Bad Request
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("id cannot be empty"))
		return
	}

	// Delete user from the database
	err := services.DeleteUser(id)
	if err != nil {

		// Check if user not found
		if err.Error() == "sql: no rows in result set" {

			// Bad Request
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("user not found"))
			fmt.Println("user not found")
		} else {

			// Internal Server Error
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
			fmt.Println("internal server error")
		}
		return
	}

	// User deleted successfully
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user deleted successfully"))
	fmt.Println("user deleted successfully")
}
