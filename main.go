package main

import (
	"log"
	"net/http"

	"github.com/golang-api-task/controllers"
)

func main() {

	http.HandleFunc("/users", controllers.CreateUser)
	http.HandleFunc("/users/", controllers.GetUsers)
	http.HandleFunc("/posts", controllers.CreatePost)
	http.HandleFunc("/posts/", controllers.GetPost)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
