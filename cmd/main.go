package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	p := Person{
		ID:        [12]byte{},
		LastName:  "Inu",
		FirstName: "Haku",
		Age:       3,
	}
	p.Age = 5

	fmt.Println("So it begins.")
	fmt.Println(p)
	r := mux.NewRouter()
	routes(r)

	// Critical to work on AppEngine
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))

}
func routes(r *mux.Router) {
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/person", personHandler)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func personHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Person!")

}
