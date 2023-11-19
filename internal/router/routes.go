package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")

	v1.GET("/opening", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "GET all openings",
		})
	})

	v1.POST("/opening", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "POST opening",
		})
	})

	v1.DELETE("/opening/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "DELETE opening",
		})
	})

	v1.PUT("/opening/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "PUT opening",
		})
	})

	v1.GET("/opening/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "GET opening",
		})
	})

}
