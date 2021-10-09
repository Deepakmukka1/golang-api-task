package main

import (
	"log"
	"net/http"

	"github.com/golang-api-task/controllers"
)

func main() {

	http.HandleFunc("/users", controllers.CreateUser)
	// http.HandleFunc("/users/:id", controllers.GetUsers)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
