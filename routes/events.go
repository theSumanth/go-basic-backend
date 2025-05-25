package routes

import (
	"net/http"
	"strconv"
	"time"

	"example.com/go-basic-backend/models"
	"example.com/go-basic-backend/utils"
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
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "could not authorize"})
		return
	}

	userMap, err := utils.VerifyToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "could not authorize", "error": err.Error()})
		return
	}

	var event models.Event

	err = ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the request", "error": err.Error()})
		return
	}

	userId, ok := userMap["userId"].(int64)
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

	_, err = models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event", "error": err.Error()})
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

	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete the event", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deletion successful"})
}
