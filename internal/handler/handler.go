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

func GetOpening(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET opening",
	})
}

func GetOpenings(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET all openings",
	})
}

func DeleteOpening(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "DELETE opening",
	})
}

func UpdateOpening(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PUT opening",
	})
}
