package routes

import (
	"net/http"

	"github.com/denizydmr07/UserManagement/backend/internal/controllers"
)

// Cors adds CORS headers to the response
// It takes a handler function as a parameter
// It returns a handler function
// It adds CORS headers to the response and calls the handler function
func Cors(next http.Handler) http.Handler {

	// Return handler function
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Add CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		next.ServeHTTP(w, r)
	})
}

// NewRouter creates a new router
// It returns a handler
// It adds routes to the router
func NewRouter() http.Handler {

	// Create a new router
	mux := http.NewServeMux()

	// Function to add a new user
	// It takes a handler function as a parameter
	// It returns a handler function
	mux.HandleFunc("/api/saveUser", func(w http.ResponseWriter, r *http.Request) {
		controllers.SaveUser(w, r)
	})

	// Function to get all users
	// It takes a handler function as a parameter
	// It returns a handler function
	mux.HandleFunc("/api/getUsers", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetUsers(w, r)
	})

	// Function to get a user with a specific ID
	// It takes a handler function as a parameter
	// It returns a handler function
	mux.HandleFunc("/api/getUser", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetUserWithID(w, r)
	})

	// Function to update a user
	// It takes a handler function as a parameter
	// It returns a handler function
	mux.HandleFunc("/api/updateUser", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateUser(w, r)
	})

	// Function to delete a user
	// It takes a handler function as a parameter
	// It returns a handler function
	mux.HandleFunc("/api/deleteUser", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteUser(w, r)
	})

	return mux
}
