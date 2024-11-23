package main

import (
	"fmt"
	"log"
	"net/http"
	"user/controller"
	"user/database"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hi Muskan...")

	// Initialize database

	db, err := database.InitDatabase()

	if err != nil {
		fmt.Errorf("Database initialiaztion failed: %v", err)
		log.Fatalf("Database initialiaztion failed: %v", err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/users", controller.CreateUser(db)).Methods("POST")
	router.HandleFunc("/api/users", controller.GetUsers(db)).Methods("GET")
	router.HandleFunc("/api/users/{id}", controller.GetUser(db)).Methods("GET")
	router.HandleFunc("/api/users/{id}", controller.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/api/users/{id}", controller.DeleteUser(db)).Methods("DELETE")

	fmt.Println("Server starting on: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
