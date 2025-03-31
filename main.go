package main

import (
	"license/auth"
	"license/config"
	"license/handlers"
	"net/http"
)

func main() {
	// Connect to db
	config.Connect()

	http.HandleFunc("/login", auth.LoginHandler)
	http.HandleFunc("/create", handlers.AddKey)
	http.HandleFunc("/delete", handlers.DeleteKey)
	http.HandleFunc("/authenticate", handlers.AuthenticateKey)

	// Listen
	http.ListenAndServe(":8090", nil)
}
