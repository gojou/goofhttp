package main

import (
	"github.com/gojou/goofhttp/pkg/httphandlers"
	"github.com/gorilla/mux"
)

// Routes defines the html paths for the app
func Routes(r *mux.Router) {
	r.HandleFunc("/", httphandlers.HomeHandler).Methods("GET")
	r.HandleFunc("/about", httphandlers.GetAboutHandler).Methods("GET")
	r.HandleFunc("/add", httphandlers.GetAddPersonHandler).Methods("GET")
	r.HandleFunc("/person", httphandlers.GetPersonHandler).Methods("GET")
	r.HandleFunc("/person/{id}", httphandlers.GetPersonHandler).Methods("GET")
	r.HandleFunc("/persons", httphandlers.GetPersonsHandler).Methods("GET")
	r.HandleFunc("/addPerson", httphandlers.AddPersonHandler).Methods("POST")

}
