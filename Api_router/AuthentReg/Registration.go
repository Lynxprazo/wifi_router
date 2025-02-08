package Registration

import (
	"Api_router/databaseconn"
	"database/sql"
	"encoding/json"
	
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type Reg struct {
	Username     string `json:"Username"`
	Password     string `json:"Password"`
	Phone_Number string `json:"Phone_Number"`
}

var DB *sql.DB

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON body into the Reg struct
	var Reg_user Reg
	err := json.NewDecoder(r.Body).Decode(&Reg_user)
	if err != nil {
		http.Error(w, "Failed to decode JSON body: "+err.Error(), http.StatusBadRequest)
		return
	}
//check if user already  exist 
 exist:= Query
	// Validate that all fields are filled
	if Reg_user.Username == "" || Reg_user.Password == "" || Reg_user.Phone_Number == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Reg_user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Insert the user into the database
	query := "INSERT INTO user_reg (Username , Password, Phone_Number) VALUES (?, ?, ?)"
	_, err = database.DB.Exec(query, Reg_user.Username, string(hashedPassword), Reg_user.Phone_Number)
	if err != nil {
		http.Error(w, "Failed to insert data into the database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a success response
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}