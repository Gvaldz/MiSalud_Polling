package controllers

import (
	"clinica/src/internal/patients/application"
	"clinica/src/internal/patients/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdatePatientsController struct {
	updatePatients *application.UpdatePatients
}

func NewUpdatePatientsController(update *application.UpdatePatients) *UpdatePatientsController{
	return &UpdatePatientsController{
		updatePatients: update,
	}
}

func (h *UpdatePatientsController) Update(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var patientRequest struct {
		Nombres      string `json:"nombres"`
		Apellidos    string `json:"apellidos"`
		Nacimiento	 string `json:"nacimiento"`
		Genero		 string `json:"genero"`
		Peso		 float32 `json:"peso"`
		Estatura	 float32 `json:"estatura"`
	}

	if err := c.ShouldBindJSON(&patientRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patient := entities.Patient{
		IdPatients:  int32(id),
		Nombres:     patientRequest.Nombres,
		Apellidos:   patientRequest.Apellidos,
		Nacimiento:  patientRequest.Nacimiento,
		Genero:      patientRequest.Genero,
		Peso:        patientRequest.Peso,
		Estatura:    patientRequest.Estatura,
	}

	err = h.updatePatients.Execute(patient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

    c.JSON(http.StatusOK, gin.H{"message": "Paciente actualizado correctamente", "paciente": patient})
}
