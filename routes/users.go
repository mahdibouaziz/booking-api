package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahdibouaziz/booking-api/models"
	"github.com/mahdibouaziz/booking-api/utils"
)

func signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create a user",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "user created",
	})
}

func login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}

	err = user.ValidateCredentails()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "logged in successfully",
		"token":   token,
	})

}
