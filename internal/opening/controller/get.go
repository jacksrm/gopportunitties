package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jacksrm/gopportunitties/internal/opening/dto"
)

// @BasePath /api/v1/

// @Summary Get Opening
// @Description Gets an opening
// @Tags openings
// @Accept json
// @Produce json
// @Param id path string true "Id of the opening to get"
// @Success 200 {object} dto.RequestResponse
// @Failure 400 {object} dto.RequestResponse
// @Router /opening/{id} [get]
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
