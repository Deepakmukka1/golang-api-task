package db

import (
	"context"
	"fmt"
	"log"

	"github.com/golang-api-task/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DatabaseInit() *mongo.Client {
	clientOptions := options.Client().ApplyURI(config.MONGO_URI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb status: Connected")
	return client
}
