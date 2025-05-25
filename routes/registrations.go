package routes

import (
	"net/http"
	"strconv"

	"example.com/go-basic-backend/models"
	"github.com/gin-gonic/gin"
)

func registerEvent(ctx *gin.Context) {
	userDetailsMap := ctx.GetStringMap("userDetailsMap")

	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the request", "error": err.Error()})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event", "error": err.Error()})
		return
	}

	userId, ok := userDetailsMap["userId"].(int64)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "invalid user id in token"})
		return
	}

	if event.IsUserRegistered(userId) {
		ctx.JSON(http.StatusConflict, gin.H{"message": "user already registered for the event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not register for the event", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "registration successful"})
}

func cancelRegister(ctx *gin.Context) {
	userDetailsMap := ctx.GetStringMap("userDetailsMap")

	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the request", "error": err.Error()})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event", "error": err.Error()})
		return
	}

	userId, ok := userDetailsMap["userId"].(int64)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "invalid user id in token"})
		return
	}

	if !event.IsUserRegistered(userId) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "user is not registered for the event"})
		return
	}

	err = event.CancelRegister(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel for the event", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "cancelling for event successful"})
}
