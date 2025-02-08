package controllers

import (
	"clinica/src/internal/patients/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllPatientsController struct {
	listPatients  *application.ListPatients
}

func NewGetAllPatientsController(list *application.ListPatients) *GetAllPatientsController {
	return &GetAllPatientsController{
		listPatients:  list,
	}
}

func (h *GetAllPatientsController) GetAll(c *gin.Context) {
	patients, err := h.listPatients.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, patients)
}