package middlewares

import (
	"net/http"

	"example.com/go-basic-backend/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "could not authorize"})
		return
	}

	userDetailsMap, err := utils.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "could not authorize", "error": err.Error()})
		return
	}

	ctx.Set("userDetailsMap", userDetailsMap)
}
