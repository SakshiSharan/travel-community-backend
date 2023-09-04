package controller

import (
	"travel-backend/constants"
	request "travel-backend/interface"
	"travel-backend/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateCommunity(ctx *gin.Context, client *mongo.Client, createCommunityRequest *request.CreateCommunity) (*model.Community, error) {
	var newId primitive.ObjectID = primitive.NewObjectID()
	community := model.Community{
		ID: newId,
		Title: createCommunityRequest.Title,
		Description: createCommunityRequest.Description,
		Members: createCommunityRequest.Members,
		Privacy: createCommunityRequest.Privacy,
		CreatedBy: createCommunityRequest.CreatedBy,
	}

	_, err := client.Database(constants.DB).Collection(constants.COLLECTION_COMMUNITIES).InsertOne(ctx, community)
	if err != nil {
		return nil, err
	}

	_, err = client.Database(constants.DB).Collection(constants.COLLECTION_USERS).UpdateMany(ctx, 
		bson.M{
			"_id": bson.M{"$in": createCommunityRequest.Members},
		},
		bson.M{
			"$push": bson.M{
				"communities": newId,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return &community, nil
}

func GetCommunity(ctx *gin.Context, client *mongo.Client, id *primitive.ObjectID) (*model.Community, error) {
	var community model.Community
	client.Database(constants.DB).Collection(constants.COLLECTION_COMMUNITIES).FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(&community)
	return &community, nil
}

func GetAllCommunities(ctx *gin.Context, client *mongo.Client) (*[]model.Community, error) {
	var communities []model.Community
	cursor, err := client.Database(constants.DB).Collection(constants.COLLECTION_COMMUNITIES).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	cursor.All(ctx, &communities)
	return &communities, nil
}

// Privacy settings later
func JoinCommunity(ctx *gin.Context, client *mongo.Client, joinCommunityRequest *request.JoinCommunity) (*model.Community, *model.User, error) {
	var user model.User
	var community model.Community

	client.Database(constants.DB).Collection(constants.COLLECTION_COMMUNITIES).FindOneAndUpdate(ctx, bson.M{
		"_id": joinCommunityRequest.CommunityID,
	},
	bson.M{
		"$push": bson.M{
			"members": joinCommunityRequest.UserID,
		},
	},
	).Decode(&community)

	client.Database(constants.DB).Collection(constants.COLLECTION_USERS).FindOneAndUpdate(ctx, bson.M{
		"_id": joinCommunityRequest.UserID,
	},
	bson.M{
		"$push": bson.M{
			"communities": joinCommunityRequest.CommunityID,
		},
	},
	).Decode(&user)

	return &community, &user, nil
}

// Search by title, location if possible