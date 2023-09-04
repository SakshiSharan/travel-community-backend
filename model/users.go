package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID   `json:"_id" bson:"_id"`
	FirstName   string               `json:"firstName" bson:"firstName"`
	LastName    string               `json:"lastName" bson:"lastName"`
	Email       string               `json:"email" bson:"email"`
	Address     string               `json:"address" bson:"address"`
	ContactNo   string               `json:"contactNo" bson:"contactNo"`
	Communities []primitive.ObjectID `json:"communities" bson:"communities"`
	Friends     []primitive.ObjectID `json:"friends" bson:"friends"`
	Trips       []primitive.ObjectID `json:"trips" bson:"trips"`
}
