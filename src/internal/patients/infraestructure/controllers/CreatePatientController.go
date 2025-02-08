package controllers

import (
	"clinica/src/internal/patients/application"
	"clinica/src/internal/patients/domain/entities"
	"net/http"
	"github.com/gin-gonic/gin"
)

type CreatePatientController struct {
	createPatient *application.CreatePatient
}

func NewCreatePatientController(create *application.CreatePatient) *CreatePatientController {
	return &CreatePatientController{
		createPatient: create,
	}
}

func (h *CreatePatientController) Create(c *gin.Context) {
	var patientRequest struct {
		Nombres     string `json:"nombres"`
		Apellidos   string `json:"apellidos"`
		Nacimiento string `json:"nacimiento"`
		Genero string `json:"genero"`
		Peso float32 `json:"peso"`
		Estatura float32 `json:"estatura"`}

	if err := c.ShouldBindJSON(&patientRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patient := entities.Patient{
		Nombres:     patientRequest.Nombres,
		Apellidos:   patientRequest.Apellidos,
		Nacimiento: patientRequest.Nacimiento,
		Genero: patientRequest.Genero,
		Peso: patientRequest.Peso,
		Estatura: patientRequest.Estatura,
	}
	
	err := h.createPatient.Execute(patient)
		if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
    c.JSON(http.StatusCreated, gin.H{"message": "Paciente creado correctamente", "paciente": patient})
}