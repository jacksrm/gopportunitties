package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOpening(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET opening",
	})
}
