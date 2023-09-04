package handler

import (
	"travel-backend/constants"
	"travel-backend/controller"
	request "travel-backend/interface"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	MongoClient *mongo.Client
}

func (h *Handler) CreateUser(ctx *gin.Context) {
	var requestBody request.CreateUser
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": constants.ERROR_INCORRECT_REQUEST,
		})
		return
	}
	user, err := controller.CreateUser(ctx, h.MongoClient, &requestBody)

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": err,
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "User created successfully",
			"user":    user,
		})
	}
}

func (h *Handler) GetUser(ctx *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(ctx.GetHeader("id"))
	user, err := controller.GetUser(ctx, h.MongoClient, &id)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": err,
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "Exams fetched successfully",
			"user":    user,
		})
	}
}

func (h *Handler) AddFriend(ctx *gin.Context) {
	var requestBody request.AddFriend
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": constants.ERROR_INCORRECT_REQUEST,
		})
		return
	}
	user, friend, err := controller.AddFriend(ctx, h.MongoClient, &requestBody)

	if err != nil {
		ctx.AbortWithStatusJSON(200, gin.H{
			"error": err,
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "Friends added successfully",
			"user":    user,
			"friend":  friend,
		})
	}
}
