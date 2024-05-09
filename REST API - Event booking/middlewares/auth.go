package middlewares

import (
	"net/http"
	"strings"

	"example.com/REST-API-Event-Booking/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	parts := strings.Split(token, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Authorization header format."})
		return
	}

	userId, err := utils.VerifyToken(parts[1])
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed."})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
