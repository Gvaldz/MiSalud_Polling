package controllers

import (
	"clinica/src/internal/patients/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetPatientCountController struct {
	useCase *application.GetPatientCount
}

func NewGetPatientCountController(useCase *application.GetPatientCount) *GetPatientCountController {
	return &GetPatientCountController{useCase: useCase}
}

func (c *GetPatientCountController) GetCount(ctx *gin.Context) {
	count, err := c.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"patients": count})
}
