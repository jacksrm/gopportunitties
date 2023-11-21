package controller

import (
	"github.com/jacksrm/gopportunitties/config"
	"github.com/jacksrm/gopportunitties/internal/opening/dto"
	"github.com/jacksrm/gopportunitties/internal/opening/service"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *service.Service
}

func New(service *service.Service) *Controller {
	return &Controller{service}
}

func sendError(context *gin.Context, status int, err error, loggerPrefix string, message string) {
	logger := config.GetLogger(loggerPrefix)
	logger.Error(message)
	context.JSON(status, dto.RequestResponse{
		Error:   true,
		Message: err.Error(),
		Data:    nil,
	})
}
