package routes

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"example.com/REST-API-Event-Booking/models"
	"github.com/gin-gonic/gin"
)

type Event struct {
	ID          int64
	Name        string     `binding:"required"`
	Description string     `binding:"required"`
	Location    string     `binding:"required"`
	DateTime    *time.Time `binding:"required"`
	UserID      int64
}

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventIdStr := context.Param("id")
	eventId, err := strconv.ParseInt(eventIdStr, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID format."})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch event."})
		return
	}

	if event == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found."})
		return
	}

	err = event.Register(userId)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register for event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered!"})
}

func cancelRegistration(context *gin.Context) {
	eventIdStr := context.Param("id")
	eventId, err := strconv.ParseInt(eventIdStr, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event registration ID format."})
		return
	}

	err = models.CancelEventRegistration(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event registration cancelled successfully."})
}
