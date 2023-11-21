package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jacksrm/gopportunitties/internal/opening"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")

	opening := opening.New()

	v1.GET("/openings", opening.Controller.GetAll)
	v1.GET("/opening/:id", opening.Controller.Get)
	v1.POST("/opening", opening.Controller.Create)
	v1.PUT("/opening/:id", opening.Controller.Update)
	v1.DELETE("/opening/:id", opening.Controller.Delete)
}
