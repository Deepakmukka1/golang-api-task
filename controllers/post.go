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
	w.Header().Set("Content-Type", "application/json")
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
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

	parts := strings.Split(r.URL.String(), "/") // parts contains the data of /post/id is split
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

	result := models.Post{}
	filter := bson.M{"_id": objectIDS}
	err := client.Database("InstagramDB").Collection("posts").FindOne(context.Background(), filter).Decode(&result)
	if err == nil {
		fmt.Println(result)
		json.NewEncoder(w).Encode(result)
	} else {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
	}

}
