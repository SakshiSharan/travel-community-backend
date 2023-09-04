package handler

import (
	"travel-backend/constants"
	"travel-backend/controller"
	request "travel-backend/interface"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) GetCommunity(ctx *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(ctx.GetHeader("id"))
	community, err := controller.GetCommunity(ctx, h.MongoClient, &id)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": err,
		})
	} else {
		ctx.JSON(200, gin.H{
			"message":   "Community fetched successfully",
			"community": community,
		})
	}
}

func (h *Handler) CreateCommunity(ctx *gin.Context) {
	var requestBody request.CreateCommunity
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": constants.ERROR_INCORRECT_REQUEST,
		})
		return
	}
	community, err := controller.CreateCommunity(ctx, h.MongoClient, &requestBody)

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": err,
		})
	} else {
		ctx.JSON(200, gin.H{
			"message":   "Community created successfully",
			"community": community,
		})
	}
}

func (h *Handler) JoinCommunity(ctx *gin.Context) {
	var requestBody request.JoinCommunity
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": constants.ERROR_INCORRECT_REQUEST,
		})
		return
	}
	community, user, err := controller.JoinCommunity(ctx, h.MongoClient, &requestBody)

	if err != nil {
		ctx.AbortWithStatusJSON(200, gin.H{
			"error": err,
		})
	} else {
		ctx.JSON(200, gin.H{
			"message":   "Community added successfully",
			"community": community,
			"user":      user,
		})
	}
}
