package request

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateUser struct {
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	Email     string `json:"email" bson:"email"`
	Address   string `json:"address" bson:"address"`
	ContactNo string `json:"contactNo" bson:"contactNo"`
}

type AddFriend struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	FriendID primitive.ObjectID `json:"friendId" bson:"friendId"`
}
