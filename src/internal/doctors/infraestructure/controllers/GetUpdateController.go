package controllers

import (
	"clinica/src/internal/doctors/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetUpdatedDoctorsController struct {
	useCase *application.GetUpdatedDoctors
}

func NewGetUpdatedDoctorsController(useCase *application.GetUpdatedDoctors) *GetUpdatedDoctorsController {
	return &GetUpdatedDoctorsController{useCase: useCase}
}

func (c *GetUpdatedDoctorsController) GetUpdatedDoctors(ctx *gin.Context) {
	since := ctx.Query("since")
	if since == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Se requiere el parámetro 'since'"})
		return
	}

	doctors, err := c.useCase.Execute(since)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, doctors)
}
