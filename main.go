package main

import (
	"context"
	"fmt"
	"os"
	"travel-backend/constants"
	"travel-backend/handler"
	"travel-backend/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	router := gin.Default()
	ctx := context.Background()

	router.Use(cors.New(utils.DefaultCors()))

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv(constants.CONNECTION_STRING)).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server-
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to MongoDB")
	}
	defer client.Disconnect(ctx)

	handler := &handler.Handler{
		MongoClient: client,
	}

	// user
	router.GET("/user/id", handler.GetUser)
	router.POST("/user", handler.CreateUser)
	router.PUT("/user/addfriend", handler.AddFriend)

	// trip
	router.GET("/trip/id", handler.GetTrip)
	router.POST("/trip", handler.CreateTrip)
	router.PUT("/trip/join", handler.JoinTrip)

	// community
	router.GET("/community/id", handler.GetCommunity)
	router.POST("/community", handler.CreateCommunity)
	router.PUT("/Community/join", handler.JoinCommunity)

	router.Run()
}
