package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-rest-api/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/health-check", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/persons", handlers.Persons).Methods("GET")

	http.Handle("/", router)

	http.ListenAndServe(":8080", router)
}
