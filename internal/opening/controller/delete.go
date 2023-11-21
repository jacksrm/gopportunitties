package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jacksrm/gopportunitties/internal/opening/dto"
)

// @BasePath /api/v1/

// @Summary Delete Opening
// @Description Deletes an opening
// @Tags openings
// @Accept json
// @Produce json
// @Param id path string true "Id of the opening to delete"
// @Success 200 {object} dto.RequestResponse
// @Failure 400 {object} dto.RequestResponse
// @Router /opening/:id [delete]
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
