package router

import (
	"github.com/jacksrm/gopportunitties/internal/opening"

	"github.com/gin-gonic/gin"
	docs "github.com/jacksrm/gopportunitties/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	const basePath = "/api/v1"
	v1 := router.Group(basePath)
	opening := opening.New()
	docs.SwaggerInfo.BasePath = basePath

	v1.GET("/openings", opening.Controller.GetAll)
	v1.GET("/opening/:id", opening.Controller.Get)
	v1.POST("/opening", opening.Controller.Create)
	v1.PUT("/opening/:id", opening.Controller.Update)
	v1.DELETE("/opening/:id", opening.Controller.Delete)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
