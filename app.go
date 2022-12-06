package main

import (
	"log"
	"net/http"
)

// define routes
func Start() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/CUSTOMERS", getAllCustomers)
	//starting server

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
