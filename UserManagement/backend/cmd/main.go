package main

import (
	"fmt"
	"net/http"

	"github.com/denizydmr07/UserManagement/backend/internal/db"
	"github.com/denizydmr07/UserManagement/backend/internal/routes"
)

func main() {
	db_err := db.InitDB()

	if db_err != nil {
		fmt.Println("DB Initialization Error")
		return
	}

	defer db.CloseDB()

	router := routes.NewRouter()

	corsRouter := routes.Cors(router)

	port := 8080
	addr := fmt.Sprintf(":%v", port)
	fmt.Printf("Server listening on port %v\n", port)

	err := http.ListenAndServe(addr, corsRouter)

	if err != nil {
		panic(err)
	}
}
