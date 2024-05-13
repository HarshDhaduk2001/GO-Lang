package routes

import (
	"net/http"

	"example.com/REST-API-Event-Booking/models"
	"example.com/REST-API-Event-Booking/utils"
	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrInvalidJSONFormat})
		return
	}

	if existingUser, err := models.FindUserByEmail(user.Email); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": utils.ErrErrorCheckingExistence})
		return
	} else if existingUser != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrUserExists})
		return
	}

	if err := user.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": utils.ErrCouldNotSaveUser})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}

func login(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrInvalidJSONFormat})
		return
	}

	if err := user.ValidateCredentials(); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": utils.ErrInvalidCredentials})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": utils.ErrFailedToGenerateToken})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login Successful!", "token": token})
}
