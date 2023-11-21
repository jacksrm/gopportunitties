package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jacksrm/gopportunitties/config"
	"github.com/jacksrm/gopportunitties/internal/opening/dto"
	"github.com/jacksrm/gopportunitties/internal/opening/service"

	"github.com/gin-gonic/gin"
)

func sendError(context *gin.Context, status int, err error, loggerPrefix string, message string) {
	logger := config.GetLogger(loggerPrefix)
	logger.Error(message)
	context.JSON(status, dto.RequestResponse{
		Error:   true,
		Message: err.Error(),
		Data:    nil,
	})
}

type Controller struct {
	service *service.Service
}

func New(service *service.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) Create(context *gin.Context) {
	const operationName = "CreateOpening"
	data := dto.CreateOpening{}

	if err := context.ShouldBindJSON(&data); err != nil {
		sendError(
			context,
			http.StatusBadRequest,
			err,
			operationName,
			fmt.Sprintf("%v: %v", operationName, err),
		)
		return
	}

	result, err := c.service.Create(data)

	if err != nil {
		sendError(
			context,
			http.StatusBadRequest,
			err,
			operationName,
			fmt.Sprintf("%v: %v", operationName, err),
		)
		return
	}

	context.JSON(http.StatusCreated, dto.RequestResponse{
		Error:   false,
		Message: "Opening created successfully",
		Data:    gin.H{"opening": result},
	})
}

func (c *Controller) Get(context *gin.Context) {
	const operationName = "GetOpening"
	id := context.Param("id")
	intId, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		sendError(
			context,
			http.StatusBadRequest,
			err,
			operationName,
			fmt.Sprintf("%v: %v", operationName, err),
		)
		return
	}

	result, err := c.service.Get(intId)
	if err != nil {
		sendError(
			context,
			http.StatusBadRequest,
			err,
			operationName,
			fmt.Sprintf("%v: %v", operationName, err),
		)
		return
	}

	context.JSON(http.StatusOK, dto.RequestResponse{
		Error:   false,
		Message: "Opening retrieved successfully",
		Data:    gin.H{"opening": result},
	})

}

func (c *Controller) GetAll(context *gin.Context) {
	const operationName = "GetOpenings"
	result, err := c.service.GetAll()

	if err != nil {
		sendError(
			context,
			http.StatusBadRequest,
			err,
			operationName,
			fmt.Sprintf("%v: %v", operationName, err),
		)
		return
	}

	context.JSON(http.StatusOK, dto.RequestResponse{
		Error:   false,
		Message: "Openings retrieved successfully",
		Data:    gin.H{"openings": result},
	})
}

func (c *Controller) Update(context *gin.Context) {
	const operationName = "UpdateOpening"
	id := context.Param("id")
	intId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		sendError(
			context,
			http.StatusBadRequest,
			err,
			operationName,
			fmt.Sprintf("%v: %v", operationName, err),
		)
		return
	}

	data := dto.UpdateOpening{}

	if err := context.ShouldBindJSON(&data); err != nil {
		sendError(
			context,
			http.StatusBadRequest,
			err,
			operationName,
			fmt.Sprintf("%v: %v", operationName, err),
		)
		return
	}

	result, err := c.service.Update(intId, data)

	if err != nil {
		sendError(
			context,
			http.StatusBadRequest,
			err,
			operationName,
			fmt.Sprintf("%v: %v", operationName, err),
		)
		return
	}

	context.JSON(http.StatusOK, dto.RequestResponse{
		Error:   false,
		Message: "Opening updated successfully",
		Data:    gin.H{"opening": result},
	})
}

func (c *Controller) Delete(context *gin.Context) {
	const operationName = "DeleteOpening"
	id := context.Param("id")
	intId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		sendError(
			context,
			http.StatusBadRequest,
			err,
			operationName,
			fmt.Sprintf("%v: %v", operationName, err),
		)
		return
	}

	result, err := c.service.Delete(intId)

	if err != nil {
		sendError(
			context,
			http.StatusBadRequest,
			err,
			operationName,
			fmt.Sprintf("%v: %v", operationName, err),
		)
		return
	}

	context.JSON(http.StatusOK, dto.RequestResponse{
		Error:   false,
		Message: "Opening deleted successfully",
		Data:    gin.H{"opening": result},
	})
}
