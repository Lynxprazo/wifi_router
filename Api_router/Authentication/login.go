package main

import "net/http"
type Login struct{
	username string`json:username`
	Password string `json:password`
}
func LoginRegister() {
	handleHttp := func(w http.ResponseWriter, r*http.Request) {

		if r.Method != "POST"{
			http.Error(w,"Methode is not allowed",http.StatusMethodNotAllowed)
			return
		}

	}

}