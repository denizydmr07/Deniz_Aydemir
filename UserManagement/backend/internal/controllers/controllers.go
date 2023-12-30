package controllers

import (
	"fmt"
	"net/http"

	"github.com/denizydmr07/UserManagement/backend/internal/services"
)

func SaveUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("username and password cannot be empty"))
		return
	}

	err := services.InsertUser(username, password)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.username" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("username already exists"))
			fmt.Println("username already exists")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
			fmt.Println("internal server error")
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user saved successfully"))
	fmt.Println("user saved successfully")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		fmt.Println("internal server error")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(users)
	fmt.Println("users retrieved successfully")
}

func GetUserWithID(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("id cannot be empty"))
		return
	}

	user, err := services.GetUserWithID(id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("user not found"))
			fmt.Println("user not found")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
			fmt.Println("internal server error")
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(user)
	fmt.Println("user retrieved successfully")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	username := r.FormValue("username")
	password := r.FormValue("password")

	if id == "" || username == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("id, username and password cannot be empty"))
		return
	}

	err := services.UpdateUser(id, username, password)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("user not found"))
			fmt.Println("user not found")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
			fmt.Println("internal server error")
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user updated successfully"))
	fmt.Println("user updated successfully")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("id cannot be empty"))
		return
	}

	err := services.DeleteUser(id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("user not found"))
			fmt.Println("user not found")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
			fmt.Println("internal server error")
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user deleted successfully"))
	fmt.Println("user deleted successfully")
}
