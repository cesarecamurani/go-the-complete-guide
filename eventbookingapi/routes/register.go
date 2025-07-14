package routes

import (
	"example.com/eventbookingapi/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetString("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), parseBase, parseBitSize)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Successfully registered for event!", "event": event})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetString("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), parseBase, parseBitSize)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
		return
	}

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel registration"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration cancelled successfully!", "event": event})
}
