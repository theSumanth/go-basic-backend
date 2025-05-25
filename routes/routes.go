package routes

import (
	"example.com/go-basic-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getSingleEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createSingleEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerEvent)
	authenticated.POST("/events/:id/cancel-register", cancelRegister)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
