package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mahdibouaziz/booking-api/models"
)

func registerEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id",
		})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event",
		})
		return
	}

	userId := ctx.GetInt64("userId")
	err = event.Register(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could register user for event",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "registered",
	})

}

func cancelRegistration(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id",
		})
		return
	}

	var event models.Event
	event.ID = eventId

	userId := ctx.GetInt64("userId")
	err = event.CancelRegistration(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could cancel registration for the event",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "registeration cancelled",
	})
}
