package controller

import (
	"errors"
	"travel-backend/constants"
	request "travel-backend/interface"
	"travel-backend/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateUser(ctx *gin.Context, client *mongo.Client, createUserRequest *request.CreateUser) (*model.User, error) {
	var newUserId primitive.ObjectID = primitive.NewObjectID()
	user := model.User{
		ID:          newUserId,
		FirstName:   createUserRequest.FirstName,
		LastName:    createUserRequest.LastName,
		Email:       createUserRequest.Email,
		Address:     createUserRequest.Address,
		ContactNo:   createUserRequest.ContactNo,
		Communities: []primitive.ObjectID{},
		Friends:     []primitive.ObjectID{},
		Trips:       []primitive.ObjectID{},
	}

	var existingUser []model.User

	cursor, err := client.Database(constants.DB).Collection(constants.COLLECTION_USERS).Find(ctx,
		bson.M{
			"email": createUserRequest.Email,
		},
	)
	if err != nil {
		return nil, err
	}
	cursor.All(ctx, &existingUser)
	defer cursor.Close(ctx)

	if len(existingUser) > 0 {
		return nil, errors.New("User already exists")
	}

	_, err = client.Database(constants.DB).Collection(constants.COLLECTION_USERS).InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(ctx *gin.Context, client *mongo.Client, id *primitive.ObjectID) (*model.User, error) {
	var user model.User
	client.Database(constants.DB).Collection(constants.COLLECTION_USERS).FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(&user)
	return &user, nil
}

func GetAllUsers(ctx *gin.Context, client *mongo.Client) (*[]model.User, error) {
	var users []model.User
	cursor, err := client.Database(constants.DB).Collection(constants.COLLECTION_USERS).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	cursor.All(ctx, &users)
	return &users, nil
}

// Add user to pending reqs of other user and add other user to sentreqs of current user
func SendFriendReq(ctx *gin.Context, client *mongo.Client, addFriendRequest *request.AddFriend) (*model.User, *model.User, error) {
	var user model.User
	var friend model.User

	client.Database(constants.DB).Collection(constants.COLLECTION_USERS).FindOneAndUpdate(ctx, bson.M{
		"_id": addFriendRequest.UserID,
	},
		bson.M{
			"$addToSet": bson.M{
				"friends": addFriendRequest.FriendID,
			},
		},
	).Decode(&user)

	client.Database(constants.DB).Collection(constants.COLLECTION_USERS).FindOneAndUpdate(ctx, bson.M{
		"_id": addFriendRequest.FriendID,
	},
		bson.M{
			"$addToSet": bson.M{
				"friends": addFriendRequest.UserID,
			},
		},
	).Decode(&friend)

	return &user, &friend, nil
}

// Very basic, later add friend request, approve and delete, different collection for friend requests or just add in user maybe.
func AddFriend(ctx *gin.Context, client *mongo.Client, addFriendRequest *request.AddFriend) (*model.User, *model.User, error) {
	var user model.User
	var friend model.User

	client.Database(constants.DB).Collection(constants.COLLECTION_USERS).FindOneAndUpdate(ctx,
		bson.M{
			"_id": addFriendRequest.UserID,
		},
		bson.M{
			"$addToSet": bson.M{
				"friends": addFriendRequest.FriendID,
			},
		},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&user)

	client.Database(constants.DB).Collection(constants.COLLECTION_USERS).FindOneAndUpdate(ctx,
		bson.M{
			"_id": addFriendRequest.FriendID,
		},
		bson.M{
			"$addToSet": bson.M{
				"friends": addFriendRequest.UserID,
			},
		},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&friend)

	return &user, &friend, nil
}
