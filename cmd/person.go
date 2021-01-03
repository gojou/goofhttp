package main

import "go.mongodb.org/mongo-driver/bson/primitive"

// Person is a Person struct
type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	LastName  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	FirstName string             `json:"firsname,omitempty" bson:"firsname,omitempty"`
	Age       int32              `json:"age,omitempty" bson:"age,omitempty"`
}
