package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jacksrm/gopportunitties/internal/opening"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")

	openingM := opening.New()

	v1.GET("/openings", openingM.Controller.GetOpenings)
	v1.POST("/opening", openingM.Controller.CreateOpening)
	v1.DELETE("/opening/:id", openingM.Controller.DeleteOpening)
	v1.PUT("/opening/:id", openingM.Controller.UpdateOpening)
	v1.GET("/opening/:id", openingM.Controller.GetOpening)
}
