package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacksrm/gopportunitties/internal/opening/dto"
)

// @BasePath /api/v1/openings
// @Description Create a new opening
// @Accept json
// @Produce json
// @Param opening body dto.CreateOpening true "Opening object that needs to be added"
// @Success 201 {object} dto.RequestResponse
// @Failure 400 {object} dto.RequestResponse
// @Router /opening [post]

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
