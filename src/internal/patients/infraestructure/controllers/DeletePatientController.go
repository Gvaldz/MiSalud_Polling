package controllers

import (
	"clinica/src/internal/patients/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeletePatientController struct {
	deletePatient *application.DeletePatient
}

func NewDeletePatientController(delete *application.DeletePatient) *DeletePatientController {
	return &DeletePatientController{
		deletePatient: delete,
	}
}

func (h *DeletePatientController) Delete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.deletePatient.Execute(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Paciente eliminado correctamente"})
}