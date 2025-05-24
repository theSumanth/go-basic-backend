package main

import (
	"net/http"
	"time"

	"example.com/go-basic-backend/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", func(ctx *gin.Context) {
		events := models.GetAllEvents()
		ctx.JSON(http.StatusOK, gin.H{"message": "successful", "events": events})
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
		event.Date = time.Now()

		event.Save()

		events := models.GetAllEvents()

		ctx.JSON(http.StatusCreated, gin.H{"message": "event created!", "events": events})
	})

	server.Run(":8080")
}
