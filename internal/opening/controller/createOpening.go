package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacksrm/gopportunitties/config"
	"github.com/jacksrm/gopportunitties/internal/opening/dto"
	"github.com/jacksrm/gopportunitties/internal/opening/repository"
	"github.com/jacksrm/gopportunitties/internal/opening/service"
)

func CreateOpening(context *gin.Context) {
	logger := config.GetLogger("CreateOpening")
	data := dto.CreateOpening{}

	if err := context.ShouldBindJSON(&data); err != nil {
		logger.Errorf("CreateOpening: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := service.CreateOpening(data, repository.SqliteOpeningRepository{})

	if err != nil {
		logger.Errorf("CreateOpening: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, result)
}
