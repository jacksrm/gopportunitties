package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteOpening(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "DELETE opening",
	})
}
