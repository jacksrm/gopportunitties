package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpening(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "POST opening",
	})
}
