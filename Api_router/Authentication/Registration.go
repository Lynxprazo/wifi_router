package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Reg struct {
	username     string `json:"username"`
	Password     string `json:"Password"`
	Phone_Number string `json:"Phone_Number"`
}

func Registration() {

	HttpRegister := func(w http.ResponseWriter, r*http.Request) {

		if r.Method != http.MethodPost{
           http.Error(w, "Error occur method used not recommended",http.StatusMethodNotAllowed)
		   return
		}
		var Reg_user Reg

		err := json.NewDecoder(r.Body).Decode(&Reg_user)

		if err !=nil{
			http.Error(w,"Failed to Decoded json body" +err.Error(),http.statusInternalServerError)
			return
		}else if Reg_user.username == nil|| Reg_user.Password == nil ||Reg_user.Phone_Number == nil{
			http.Error(w,"User should fill all the text box", http.StatusBadRequest)
			return
		} 
		sql,err := query("INSERT INTO TABLE user_reg (Username, Password Pnone_Number) value(re)")






	}

}