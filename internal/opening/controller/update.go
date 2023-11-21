package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jacksrm/gopportunitties/internal/opening/dto"
)

// @BasePath /api/v1/

// @Summary Update Opening
// @Description Updates a opening
// @Tags openings
// @Accept json
// @Produce json
// @Param request body dto.UpdateOpening true "Data needed to update an opening"
// @Param id path string true "Id of the opening to update"
// @Success 200 {object} dto.RequestResponse
// @Failure 400 {object} dto.RequestResponse
// @Router /opening/{id} [put]
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
