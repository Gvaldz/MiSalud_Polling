package controllers

import (
	"clinica/src/internal/doctors/application"
	"clinica/src/internal/doctors/domain/entities"
	"net/http"
	"github.com/gin-gonic/gin"
)

type CreateDoctorController struct {
	createDoctor *application.CreateDoctor
}

func NewCreateDoctorController(create *application.CreateDoctor) *CreateDoctorController {
	return &CreateDoctorController{
		createDoctor: create,
	}
}

func (h *CreateDoctorController) Create(c *gin.Context) {
	var doctorRequest struct {
		Cedula      int32  `json:"cedula"`
		Nombres     string `json:"nombres"`
		Apellidos   string `json:"apellidos"`
		Especialidad string `json:"especialidad"`
	}

	if err := c.ShouldBindJSON(&doctorRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	doctor := entities.Doctor{
		Cedula:      doctorRequest.Cedula,
		Nombres:     doctorRequest.Nombres,
		Apellidos:   doctorRequest.Apellidos,
		Especialidad: doctorRequest.Especialidad,
	}
	
	err := h.createDoctor.Execute(doctor)
		if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
    c.JSON(http.StatusCreated, gin.H{"message": "Doctor creado correctamente", "doctor": doctor})
}