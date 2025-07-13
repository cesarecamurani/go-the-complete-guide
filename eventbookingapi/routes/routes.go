package routes

import (
	"example.com/eventbookingapi/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Public routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/signup", signUp)
	server.POST("/login", login)

	// Protected routes
	protected := server.Group("/")
	protected.Use(middlewares.Authenticate)

	protected.POST("/events", createEvent)
	protected.PUT("/events/:id", updateEvent)
	protected.DELETE("/events/:id", deleteEvent)
}
