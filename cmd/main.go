package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("So it begins.")
	r := mux.NewRouter()
	Routes(r)

	// Critical to work on AppEngine
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))

}

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello World!")
// }

// func getPersonHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("content-type", "application/json")
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
// 		"mongodb+srv://mongouser:hpkns372@cluster0.dfch4.mongodb.net/test?retryWrites=true&w=majority",
// 	))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	goofDB := client.Database("test")
// 	goofCollection := goofDB.Collection("persons")

// 	params := mux.Vars(r)
// 	id, err := primitive.ObjectIDFromHex(params["id"])
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var person Person
// 	err = goofCollection.FindOne(ctx, Person{ID: id}).Decode(&person)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(`{ "message" : "` + err.Error() + `"}`))
// 		return
// 	}
// 	json.NewEncoder(w).Encode(person)

// }

// func getPersonsHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("content-type", "application/json")
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
// 		"mongodb+srv://mongouser:hpkns372@cluster0.dfch4.mongodb.net/test?retryWrites=true&w=majority",
// 	))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	goofDB := client.Database("test")
// 	goofCollection := goofDB.Collection("persons")

// 	cursor, err := goofCollection.Find(ctx, bson.M{})
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(`{ "message" : "` + err.Error() + `"}`))
// 		return
// 	}
// 	defer cursor.Close(ctx)
// 	var persons []Person
// 	for cursor.Next(ctx) {
// 		var person Person
// 		cursor.Decode(&person)
// 		persons = append(persons, person)
// 	}
// 	if err = cursor.Err(); err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(`{ "message" : "` + err.Error() + `"}`))
// 		return
// 	}
// 	json.NewEncoder(w).Encode(persons)

// }

// func addPersonHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("content-type", "application/json")
// 	var person Person
// 	json.NewDecoder(r.Body).Decode(&person)

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
// 		"mongodb+srv://mongouser:hpkns372@cluster0.dfch4.mongodb.net/test?retryWrites=true&w=majority",
// 	))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	goofDB := client.Database("test")
// 	goofCollection := goofDB.Collection("persons")

// 	result, err := goofCollection.InsertOne(ctx, person)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	json.NewEncoder(w).Encode(result)

// }
// func getAboutHandler(w http.ResponseWriter, r *http.Request) {

// 	page := template.Must(template.ParseFiles(
// 		"static/html/_base.html",
// 		"static/html/about.html",
// 	))

// 	page.Execute(w, nil)

// }

// func getAddPersonHandler(w http.ResponseWriter, r *http.Request) {
// 	page := template.Must(template.ParseFiles(
// 		"static/html/_base.html",
// 		"static/html/addperson.html",
// 	))

// 	page.Execute(w, nil)

// }
