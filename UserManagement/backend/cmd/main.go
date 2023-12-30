package main

import (
	"fmt"
	"net/http"

	"github.com/denizydmr07/UserManagement/backend/internal/db"
	"github.com/denizydmr07/UserManagement/backend/internal/routes"
)

func main() {
	// Initialize DB
	db_err := db.InitDB()

	// Check DB Initialization
	if db_err != nil {
		fmt.Println("DB Initialization Error")
		return
	}

	// Close DB when main function ends
	defer db.CloseDB()

	// Initialize Router
	router := routes.NewRouter()

	// Initialize CORS Router
	corsRouter := routes.Cors(router)

	// Start Server
	port := 8080
	addr := fmt.Sprintf(":%v", port)
	fmt.Printf("Server listening on port %v\n", port)

	// Listen and Serve
	err := http.ListenAndServe(addr, corsRouter)

	// Check Listen and Serve Error
	if err != nil {
		panic(err)
	}
}
