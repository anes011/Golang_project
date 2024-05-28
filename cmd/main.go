package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anes011/gop/pkg/database"
	"github.com/anes011/gop/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	db, err := database.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router.HandleFunc("/users/{offset}/{limit}", handlers.HandleGetUsers(db))
	// router.HandleFunc("/userById/{id}", handlers.HandleUserById(db))
	// router.HandleFunc("/createUser", handlers.HandleCreateUser(db))
	// router.HandleFunc("/books", handlers.HandleGetBooks(db))

	fmt.Printf("Server running on port: %v", 8000)

	log.Fatal(http.ListenAndServe(":8000", router))
}
