package main

import (
	"example.com/REST-API-Event-Booking/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.POST(`/signup`, createUser)

	server.Run(":8000") // localhost:8000
}
