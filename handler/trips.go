package handler

import (
	"travel-backend/constants"
	"travel-backend/controller"
	request "travel-backend/interface"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) GetTrip(ctx *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(ctx.GetHeader("id"))
	trip, err := controller.GetTrip(ctx, h.MongoClient, &id)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": err,
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "Trip fetched successfully",
			"trip":    trip,
		})
	}
}

func (h *Handler) CreateTrip(ctx *gin.Context) {
	var requestBody request.CreateTrip
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": constants.ERROR_INCORRECT_REQUEST,
		})
		return
	}
	trip, err := controller.CreateTrip(ctx, h.MongoClient, &requestBody)

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": err,
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "Trip created successfully",
			"user":    trip,
		})
	}
}

func (h *Handler) JoinTrip(ctx *gin.Context) {
	var requestBody request.JoinTrip
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": constants.ERROR_INCORRECT_REQUEST,
		})
		return
	}
	trip, user, err := controller.JoinTrip(ctx, h.MongoClient, &requestBody)

	if err != nil {
		ctx.AbortWithStatusJSON(200, gin.H{
			"error": err,
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "Trip added successfully",
			"trip":    trip,
			"user":    user,
		})
	}
}
