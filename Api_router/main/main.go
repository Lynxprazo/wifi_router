package main

import (

	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
r :=mux.NewRouter();

r.HandleFunc("/",HandleFormRegister).Method("POST")
Cors := handlers.CORS{
handlers.AllowedOrigins([] string{any}),
handlers.AllowedMethods([]string("POST", "GET", "PUT" , "DELETE")),
handlers.AllowedHeaders([]string("content-type", "Authorization")),

}

fatal.log(http.ListenAndServe(":8080",Cors))


}