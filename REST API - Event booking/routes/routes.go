package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST(`/signup`, createUser)
	server.POST(`/login`, login)

	// Events
	server.POST(`/events`, createEvent)
}
