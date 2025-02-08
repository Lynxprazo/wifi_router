package _handlers


import (
	"Api_router/databaseconn"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB
type Reg struct {
	Username     string `json:"Username"`
	Password     string `json:"Password"`
	Phone_Number string `json:"Phone_Number"`
}
type Log struct {
	Username string `json:"username"`
	Password string `json:"Password"`
}

func LoginRegister(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Request methode isn`t allowed ", http.StatusBadRequest)
		return
	}
	var userlog Log
	err := json.NewDecoder(r.Body).Decode(&userlog)
	if err != nil {
		fmt.Print("Failed to Decode the json file", err)
	}
	query := "SELECT Password FROM user_reg WHERE username = ? "
	var hashpassword string
	err = database.DB.QueryRow(query, userlog.Username).Scan(&hashpassword)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "The Username is not found ", http.StatusUnauthorized)
		} else {
			http.Error(w, "Database Error ", http.StatusInternalServerError)

		}
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(userlog.Password))
	if err != nil {
		http.Error(w, "Invalid Password ", http.StatusUnauthorized)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfuly login"))
}





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
 var existPhone_Number string

   checkquery := "SELECT Phone_Number FROM user_reg WHERE Phone_Number = ? "

   err = database.DB.QueryRow(checkquery, Reg_user.Phone_Number).Scan(&existPhone_Number)
   if err == nil{
    w.WriteHeader(http.StatusConflict)
	json.NewEncoder(w).Encode(map[string]string{"error":"Phone number already exist"})
	return
   }else if err != sql.ErrNoRows{
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"error":"Failed to connect to the database: "+ err.Error()})
	return
   }
 

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