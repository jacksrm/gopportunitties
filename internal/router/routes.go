package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jacksrm/gopportunitties/internal/opening/controller"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")
	controller.Init()

	v1.GET("/openings", controller.GetOpenings)
	v1.POST("/opening", controller.CreateOpening)
	v1.DELETE("/opening/:id", controller.DeleteOpening)
	v1.PUT("/opening/:id", controller.UpdateOpening)
	v1.GET("/opening/:id", controller.GetOpening)

}
