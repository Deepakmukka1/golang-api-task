package main

import (
	"log"
	"net/http"

	"github.com/golang-api-task/controllers"
)

func main() {
	http.HandleFunc("/createProfile", controllers.CreateProfile)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
