package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-api-task/db"
	"github.com/golang-api-task/models"
)

func CreateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // for adding       //Content-type
	var person models.User
	err := json.NewDecoder(r.Body).Decode(&person) // storing in person   //variable of type user
	if err != nil {
		fmt.Print(err)
	}
	var userCollection = db.Database().Database("goTest").Collection("users")
	insertResult, err := userCollection.InsertOne(context.TODO(), person)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(insertResult.InsertedID) // return the //mongodb ID of generated document
}
