package request

import (
	"travel-backend/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateTrip struct {
	Title         string               `json:"title" bson:"title"`
	Description   string               `json:"description" bson:"description"`
	Size          int32                `json:"size" bson:"size"`
	Members       []primitive.ObjectID `json:"members" bson:"members"`
	CreatedBy     primitive.ObjectID   `json:"createdBy" bson:"createdBy"`
	Type          string               `json:"type" bson:"type"`
	StartLocation model.Location       `json:"startLocation" bson:"startLocation"`
	EndLocation   model.Location       `json:"endLocation" bson:"endLocation"`
	Departure     string               `json:"departure" bson:"departure"`
	Arrival       string               `json:"arrival" bson:"arrival"`
	Status        string               `json:"status" bson:"status"`
	Privacy       string               `json:"privacy" bson:"privacy"`
}

type JoinTrip struct {
	UserID primitive.ObjectID `json:"userId" bson:"userId"`
	TripID primitive.ObjectID `json:"tripId" bson:"tripId"`
}
