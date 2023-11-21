package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacksrm/gopportunitties/internal/opening/dto"
)

// @BasePath /api/v1/

// @Summary Get All Openings
// @Description Gets all openings
// @Tags openings
// @Accept json
// @Produce json
// @Success 200 {object} dto.RequestResponse
// @Failure 400 {object} dto.RequestResponse
// @Router /openings [get]
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
