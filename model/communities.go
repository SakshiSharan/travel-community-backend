package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Community struct {
	ID          primitive.ObjectID   `json:"_id" bson:"_id"`
	Title       string               `json:"title" bson:"title"`
	Description string               `json:"description" bson:"description"`
	Members     []primitive.ObjectID `json:"members" bson:"members"`
	Privacy     string               `json:"privacy" bson:"privacy"`
}
