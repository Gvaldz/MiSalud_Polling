package controllers

import (
	"clinica/src/internal/doctors/application"
	"clinica/src/internal/doctors/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateDoctorController struct {
	updateDoctor *application.UpdateDoctor
}

func NewUpdateDoctorController(update *application.UpdateDoctor) *UpdateDoctorController{
	return &UpdateDoctorController{
		updateDoctor: update,
	}
}

func (h *UpdateDoctorController) Update(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var doctorRequest struct {
		Cedula       int32  `json:"cedula"`
		Nombres      string `json:"nombres"`
		Apellidos    string `json:"apellidos"`
		Especialidad string `json:"especialidad"`
	}

	if err := c.ShouldBindJSON(&doctorRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	doctor := entities.Doctor{
		IdDoctores:  int32(id),
		Cedula: 	doctorRequest.Cedula,
		Nombres:     doctorRequest.Nombres,
		Apellidos:   doctorRequest.Apellidos,
		Especialidad: doctorRequest.Especialidad,
	}

	err = h.updateDoctor.Execute(doctor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

    c.JSON(http.StatusOK, gin.H{"message": "Doctor actualizado correctamente", "doctor": doctor})
}
