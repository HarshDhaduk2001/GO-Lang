package routes

import (
	"example.com/REST-API-Event-Booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST(`/signup`, createUser)
	server.POST(`/login`, login)

	// Events
	server.GET(`/events`, getEvents)
	server.POST(`/events`, middlewares.Authenticate, createEvent)
}
