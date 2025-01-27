package main

import (
	"net/http"
	"license/config"
	"license/handlers"
)

func main() {
	// Connect to db
	config.Connect()

	http.HandleFunc("/create", handlers.AddKey)

	// Listen	
	http.ListenAndServe(":8090", nil)
}