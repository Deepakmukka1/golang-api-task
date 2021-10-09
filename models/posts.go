package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	Caption   string             `json:"caption"`
	ImageURL  string             `json:"imageurl"`
	TimeStamp primitive.DateTime `json:"timestamp"`
}
