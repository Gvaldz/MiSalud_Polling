package controllers

import (
	"clinica/src/internal/patients/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetPatientByIdController struct {
	getPatientById *application.GetPatientById
}

func NewGetPatientByIdController(getPatientById *application.GetPatientById) *GetPatientByIdController {
	return &GetPatientByIdController{getPatientById: getPatientById}
}

func (h *GetPatientByIdController) GetById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	patient, err := h.getPatientById.Execute(int32(idInt))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, patient)
}