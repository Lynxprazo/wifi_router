package main

import (

	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"Api_router/Authentication"
)

func main() {
r :=mux.NewRouter();

r.HandleFunc("/",Registration.Registration()).Method("POST")
Cors := handlers.CORS{
handlers.AllowedOrigins([] string{any}),
handlers.AllowedMethods([]string("POST", "GET", "PUT" , "DELETE")),
handlers.AllowedHeaders([]string("content-type", "Authorization")),

}

log.fatal(http.ListenAndServe(":8080",Cors))


}