package main

import (
	"net/http"
	"strconv"
	"time"

	"example.com/go-basic-backend/db"
	"example.com/go-basic-backend/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	server.GET("/events", func(ctx *gin.Context) {
		events, err := models.GetAllEvents()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch rows from db", "error": err})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "successful", "events": events})
	})

	server.GET("/events/:id", func(ctx *gin.Context) {
		eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the event id"})
			return
		}

		event, err := models.GetEventByID(eventId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "successfully fetched the event", "event": event})
	})

	server.POST("/events", func(ctx *gin.Context) {
		var event models.Event

		err := ctx.ShouldBindJSON(&event)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the request", "error": err})
			return
		}

		event.ID = 1
		event.UserID = 1
		event.DateTime = time.Now()

		err = event.Save()

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not create the event", "error": err})
			return
		}

		events, err := models.GetAllEvents()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch rows from db", "error": err})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "event created!", "events": events})
	})

	server.Run(":8080")
}
