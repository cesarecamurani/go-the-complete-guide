package routes

import (
	"example.com/eventbookingapi/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// Protected routes
	protected := server.Group("/")
	protected.Use(middlewares.Authenticate)
	protected.POST("/events", createEvent)
	protected.PUT("/events/:id", updateEvent)
	protected.DELETE("/events/:id", deleteEvent)
	protected.POST("/events/:id/register", registerForEvent)
	protected.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signUp)
	server.POST("/login", login)
}
