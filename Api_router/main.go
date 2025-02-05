package main

import (
	"Api_router/AuthentLog" // Import the Login package
	"Api_router/AuthentReg" // Import the Registration package
	"Api_router/_handlers"
	"Api_router/databaseconn" // Import the database package
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database connection
	database.InitDB()

	// Pass the database connection to the Registration and Login packages
	Registration.DB = database.DB
	Login.DB = database.DB

	// Create a new Gorilla Mux router
	r := mux.NewRouter()

	// Register the HTTP handler for user registration and login
	r.HandleFunc("/register", Registration.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", _handlers.LoginRegister).Methods("POST")

	// Configure CORS middleware
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Allow all origins
		handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	// Start the HTTP server with CORS middleware
	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", cors(r)); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}