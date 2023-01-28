package handlers

import (
	"encoding/json"
	"fmt"
	"go-rest-api/model"
	"log"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Print("Entering health check")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

func Persons(w http.ResponseWriter, r *http.Request) {
	log.Println("entering persons end point")
	var response model.Response
	persons := prepareResponse()

	response.Persons = persons

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println("Error during json marshalling")
	}
	w.Write(jsonResponse)
}

func prepareResponse() []model.Person {
	var persons []model.Person

	var person model.Person
	person.Id = 1
	person.FirstName = "Issac"
	person.LastName = "N"
	persons = append(persons, person)

	person.Id = 2
	person.FirstName = "Albert"
	person.LastName = "E"
	persons = append(persons, person)

	person.Id = 3
	person.FirstName = "Thomas"
	person.LastName = "E"
	persons = append(persons, person)
	return persons
}
