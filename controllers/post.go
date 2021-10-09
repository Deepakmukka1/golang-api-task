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
)

func CreatePost(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		fmt.Fprint(w, "You can only do POST on this route")
		return
	}
	client := db.DatabaseInit()
	w.Header().Set("Content-Type", "application/json") // for adding       //Content-type
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post) // storing in person   //variable of type user
	if err != nil {
		fmt.Print(err)
	}
	post.Id = primitive.NewObjectID()
	var postCollection = client.Database("InstagramDB").Collection("posts")
	insertResult, err := postCollection.InsertOne(context.TODO(), post)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Post created successfully with  %s", insertResult.InsertedID)
}

func GetPost(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		fmt.Fprint(w, "You can only do GET on this route")
		return
	}
	client := db.DatabaseInit()
	w.Header().Set("Content-Type", "application/json") // for adding       //Content-type
	stringID := "ObjectId" + parts[2] + ")"
	SingleResult := client.Database("InstagramDB").Collection("users").FindOne(context.TODO(), bson.M{"_id": stringID})
	// result.Decode(models.User{})
	// jsonValue, _ := json.Marshal(SingleResult)
	// fmt.Fprintf(w, parts[2])
	fmt.Fprint(w, SingleResult.Decode(models.User{}))
	// json.NewEncoder(w).Encode(result)

}
