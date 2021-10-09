package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-api-task/db"
	"github.com/golang-api-task/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		fmt.Fprint(w, "You can only do POST on this route")
		return
	}
	client := db.DatabaseInit()
	w.Header().Set("Content-Type", "application/json")
	var person models.User
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		fmt.Print(err)
	}
	person.Id = primitive.NewObjectID()
	hashedPassword, err := HashPassword(person.Password)
	if err == nil {
		person.Password = hashedPassword
	}
	var userCollection = client.Database("InstagramDB").Collection("users")
	insertResult, err := userCollection.InsertOne(context.TODO(), person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "User created successfully with  %s", insertResult.InsertedID)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.String(), "/") // parts contains the data of /user/id is split
	if len(parts) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		fmt.Fprint(w, "You can only do GET on this route")
		return
	}
	client := db.DatabaseInit()

	w.Header().Set("Content-Type", "application/json")

	objectIDS, _ := primitive.ObjectIDFromHex(parts[2])

	result := models.User{}
	filter := bson.M{"_id": objectIDS}
	err := client.Database("InstagramDB").Collection("users").FindOne(context.Background(), filter).Decode(&result)
	if err == nil {
		fmt.Println(result)
		json.NewEncoder(w).Encode(result)
	} else {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
	}

}
