package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Location struct {
	Latitude  string `json:"latitude" bson:"latitude"`
	Longitude string `json:"longitude" bson:"longitude"`
}

type Trip struct {
	ID            primitive.ObjectID   `json:"_id" bson:"_id"`
	Title         string               `json:"title" bson:"title"`
	Description   string               `json:"description" bson:"description"`
	Size          int32                `json:"size" bson:"size"`
	Members       []primitive.ObjectID `json:"members" bson:"members"`
	CreatedBy     primitive.ObjectID   `json:"createdBy" bson:"createdBy"`
	Type          string               `json:"type" bson:"type"`
	StartLocation Location             `json:"startLocation" bson:"startLocation"`
	EndLocation   Location             `json:"endLocation" bson:"endLocation"`
	Departure     string               `json:"departure" bson:"departure"`
	Arrival       string               `json:"arrival" bson:"arrival"`
	Status        string               `json:"status" bson:"status"`
	Privacy       string               `json:"privacy" bson:"privacy"`
}
