package routes

import (
	"net/http"

	"example.com/go-basic-backend/models"
	"example.com/go-basic-backend/utils"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the request.", "error": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not create user.", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "creation of user successful."})
}

func login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the request.", "error": err.Error()})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not authenticate the user.", "error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.UserID, user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not authenticate the user.", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "login successful", "token": token})
}
