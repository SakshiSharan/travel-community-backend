package request

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateCommunity struct {
	Title       string               `json:"title" bson:"title"`
	Description string               `json:"description" bson:"description"`
	Members     []primitive.ObjectID `json:"members" bson:"members"`
	Privacy     string               `json:"privacy" bson:"privacy"`
	CreatedBy   primitive.ObjectID   `json:"createdBy" bson:"createdBy"`
}

type JoinCommunity struct {
	UserID      primitive.ObjectID `json:"userId" bson:"userId"`
	CommunityID primitive.ObjectID `json:"communityId" bson:"communityId"`
}
