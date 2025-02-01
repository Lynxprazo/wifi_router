package Registration

import (
	"Api_router"
	"database/sql"
	"encoding/json"
	"net/http"
)

type Reg struct {
	Username     string `json:"username"`
	Password     string `json:"Password"`
	Phone_Number string `json:"Phone_Number"`
}
 var DB*sql.DB
func RegisterHandler(w http.ResponseWriter , r*http.Request) {

		if r.Method != http.MethodPost{
           http.Error(w, "Error occur method used not recommended",http.StatusMethodNotAllowed)
		   return
		}
		var Reg_user Reg

		err := json.NewDecoder(r.Body).Decode(&Reg_user)

		if err !=nil{
			http.Error(w,"Failed to Decoded json body" +err.Error(),http.StatusInternalServerError)
			return
		}else if Reg_user.Username == "" || Reg_user.Password == "" ||Reg_user.Phone_Number ==  ""{
			http.Error(w,"User should fill all the text box", http.StatusBadRequest)
			return
		} 
		query :="INSERT INTO TABLE user_reg (Username, Password Phnone_Number) value(?,?,?)"

		_,err = database.DB.Exec(query,Reg_user.Username,Reg_user.Password,Reg_user.Phone_Number)
		if err != nil{
			http.Error(w , "Failed to Insert data to the database", http.StatusBadRequest)
			return

		}

}