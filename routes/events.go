package routes

import (
	"net/http"
	"strconv"
	"time"

	"example.com/go-basic-backend/models"
	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch rows from db", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successful", "events": events})
}

func getSingleEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the event id", "error": err.Error()})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully fetched the event", "event": event})
}

func createSingleEvent(ctx *gin.Context) {
	userDetailsMap := ctx.GetStringMap("userDetailsMap")

	var event models.Event

	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the request", "error": err.Error()})
		return
	}

	userId, ok := userDetailsMap["userId"].(int64)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "invalid user id in token"})
		return
	}

	event.UserID = userId
	event.DateTime = time.Now()

	err = event.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not create the event", "error": err.Error()})
		return
	}

	// events, err := models.GetAllEvents()
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch rows from db", "error": err})
	// 	return
	// }

	ctx.JSON(http.StatusCreated, gin.H{"message": "event created!", "event": event})
}

func updateEvent(ctx *gin.Context) {
	userDetailsMap := ctx.GetStringMap("userDetailsMap")

	var event models.Event

	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the requst", "error": err.Error()})
		return
	}

	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse the requst", "error": err.Error()})
		return
	}

	fetchedEvent, err := models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event", "error": err.Error()})
		return
	}

	userId, ok := userDetailsMap["userId"].(int64)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "invalid user id in token"})
		return
	}

	if fetchedEvent.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to update event"})
		return
	}

	event.ID = eventId
	err = event.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not update the event", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updating event successful!", "event": event})
}

func deleteEvent(ctx *gin.Context) {
	userDetailsMap := ctx.GetStringMap("userDetailsMap")

	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the event id", "error": err.Error()})
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

	if event.ID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to delete event"})
		return
	}

	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete the event", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deletion successful"})
}
