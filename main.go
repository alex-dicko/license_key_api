package main

import (
	"license/config"
	"license/handlers"
	"net/http"
)

func main() {
	// Connect to db
	config.Connect()

	http.HandleFunc("/create", handlers.AddKey)
	http.HandleFunc("/delete", handlers.DeleteKey)

	// Listen
	http.ListenAndServe(":8090", nil)
}
