package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOpenings(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET all openings",
	})
}
