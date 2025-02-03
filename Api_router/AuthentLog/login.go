package Login

import (
	"Api_router/databaseconn"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

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
