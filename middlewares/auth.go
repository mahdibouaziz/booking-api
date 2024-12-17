package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahdibouaziz/booking-api/utils"
)

func Authenticate(ctx *gin.Context) {
	// Token validation
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
		})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
		})
		return
	}

	ctx.Set("userId", userId)
	ctx.Next()
}
