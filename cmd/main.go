package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Person is a Person struct
type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Age       int32              `json:"age,omitempty" bson:"age,omitempty"`
}

func main() {

	fmt.Println("So it begins.")
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
	r.HandleFunc("/persons", personsHandler)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func personHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Person!")

}

func personsHandler(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://mongouser:hpkns372@cluster0.dfch4.mongodb.net/test?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}
	goofDB := client.Database("test")
	goofCollection := goofDB.Collection("persons")

	cursor, err := goofCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var persons []Person
	if err = cursor.All(ctx, &persons); err != nil {
		log.Fatal(err)
	}
	for _, person := range persons {
		fmt.Fprintf(w, "%v\n", person.ID)
		fmt.Fprintf(w, "\t%v\n", person.FirstName)
		fmt.Fprintf(w, "\t%v\n", person.LastName)
		fmt.Fprintf(w, "\t%v\n", person.Age)
	}
}
