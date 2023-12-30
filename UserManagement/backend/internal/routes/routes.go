package routes

import (
	"net/http"

	"github.com/denizydmr07/UserManagement/backend/internal/controllers"
)

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		next.ServeHTTP(w, r)
	})
}

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Function to add a new user
	mux.HandleFunc("/saveUser", func(w http.ResponseWriter, r *http.Request) {
		controllers.SaveUser(w, r)
	})

	// Function to get all users
	mux.HandleFunc("/getUsers", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetUsers(w, r)
	})

	// Function to get a user with a specific ID
	mux.HandleFunc("/getUser", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetUserWithID(w, r)
	})

	mux.HandleFunc("/updateUser", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateUser(w, r)
	})

	mux.HandleFunc("/deleteUser", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteUser(w, r)
	})

	return mux
}
