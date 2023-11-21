package controller

import (
	"net/http"

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

func (c *Controller) CreateOpening(context *gin.Context) {
	logger := config.GetLogger("CreateOpening")
	data := dto.CreateOpening{}

	if err := context.ShouldBindJSON(&data); err != nil {
		logger.Errorf("CreateOpening: %v", err)
		context.JSON(http.StatusBadRequest, dto.RequestResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	result, err := c.service.CreateOpening(data)

	if err != nil {
		logger.Errorf("CreateOpening: %v", err)
		context.JSON(http.StatusBadRequest, dto.RequestResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	context.JSON(http.StatusCreated, dto.RequestResponse{
		Error:   false,
		Message: "Opening created successfully",
		Data:    gin.H{"opening": result},
	})
}

func (c *Controller) GetOpening(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "DELETE opening",
	})
}

func (c *Controller) GetOpenings(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET all openings",
	})
}

func (c *Controller) UpdateOpening(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "UPDATE opening",
	})
}

func (c *Controller) DeleteOpening(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "DELETE opening",
	})
}
