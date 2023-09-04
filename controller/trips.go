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

func CreateTrip(ctx *gin.Context, client *mongo.Client, createTripRequest *request.CreateTrip) (*model.Trip, error) {
	var newId primitive.ObjectID = primitive.NewObjectID()
	trip := model.Trip{
		ID: newId,
		Title: createTripRequest.Title,
		Description: createTripRequest.Description,
		Size: createTripRequest.Size,
		Members: createTripRequest.Members,
		CreatedBy: createTripRequest.CreatedBy,
		Type: createTripRequest.Type,
		StartLocation: createTripRequest.StartLocation,
		EndLocation: createTripRequest.EndLocation,
		Departure: createTripRequest.Departure,
		Arrival: createTripRequest.Arrival,
		Status: createTripRequest.Status,
		Privacy: createTripRequest.Privacy,
	}

	_, err := client.Database(constants.DB).Collection(constants.COLLECTION_TRIPS).InsertOne(ctx, trip)
	if err != nil {
		return nil, err
	}

	_, err = client.Database(constants.DB).Collection(constants.COLLECTION_USERS).UpdateMany(ctx, 
		bson.M{
			"_id": bson.M{"$in": createTripRequest.Members},
		},
		bson.M{
			"$addToSet": bson.M{
				"trips": newId,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return &trip, nil
}

func GetTrip(ctx *gin.Context, client *mongo.Client, id *primitive.ObjectID) (*model.Trip, error) {
	var trip model.Trip
	client.Database(constants.DB).Collection(constants.COLLECTION_TRIPS).FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(&trip)
	return &trip, nil
}

func GetAllTrips(ctx *gin.Context, client *mongo.Client) (*[]model.Trip, error) {
	var trips []model.Trip
	cursor, err := client.Database(constants.DB).Collection(constants.COLLECTION_TRIPS).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	cursor.All(ctx, &trips)
	return &trips, nil
}

// Search
// func ShowTrips(ctx *gin.Context, client *mongo.Client) {
	
// }

// Add privacy functionalities later on
func JoinTrip(ctx *gin.Context, client *mongo.Client, joinTripRequest *request.JoinTrip) (*model.Trip, *model.User, error) {
	var user model.User
	var trip model.Trip

	client.Database(constants.DB).Collection(constants.COLLECTION_TRIPS).FindOneAndUpdate(ctx, bson.M{
		"_id": joinTripRequest.TripID,
	},
	bson.M{
		"$addToSet": bson.M{
			"members": joinTripRequest.UserID,
		},
	},
	).Decode(&trip)

	client.Database(constants.DB).Collection(constants.COLLECTION_USERS).FindOneAndUpdate(ctx, bson.M{
		"_id": joinTripRequest.UserID,
	},
	bson.M{
		"$addToSet": bson.M{
			"trips": joinTripRequest.TripID,
		},
	},
	).Decode(&user)

	return &trip, &user, nil
}