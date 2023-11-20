package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOpening(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PUT opening",
	})
}
