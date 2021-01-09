package httphandlers

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gojou/goofhttp/pkg/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// HomeHandler shows the Home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

// GetPersonHandler displays data for a single person, matched by database ID
func GetPersonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
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

	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	var person models.Person
	err = goofCollection.FindOne(ctx, models.Person{ID: id}).Decode(&person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message" : "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(w).Encode(&person)

}

// GetPersonsHandler displays all Persons
func GetPersonsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message" : "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	var persons []models.Person
	for cursor.Next(ctx) {
		var person models.Person
		cursor.Decode(&person)
		persons = append(persons, person)
	}
	if err = cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message" : "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(w).Encode(persons)

}

// AddPersonHandler adds a Person
func AddPersonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var person models.Person
	json.NewDecoder(r.Body).Decode(&person)

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

	result, err := goofCollection.InsertOne(ctx, person)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(result)

}

// GetAddPersonHandler will add a Person
func GetAddPersonHandler(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/addperson.html",
	))

	page.Execute(w, nil)

}
