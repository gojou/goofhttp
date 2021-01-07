package main

import "github.com/gorilla/mux"

// Routes defines the html paths for the app
func Routes(r *mux.Router) {
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/about", getAboutHandler).Methods("GET")
	r.HandleFunc("/person", getPersonHandler).Methods("GET")
	r.HandleFunc("/person/{id}", getPersonHandler).Methods("GET")
	r.HandleFunc("/persons", getPersonsHandler).Methods("GET")
	r.HandleFunc("/addPerson", addPersonHandler).Methods("POST")

}
