package controllers

import (
	"clinica/src/internal/doctors/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetDoctorByIdController struct {
	getDoctorById *application.GetDoctorById
}

func NewGetDoctorByIdController(getDoctorById *application.GetDoctorById) *GetDoctorByIdController {
	return &GetDoctorByIdController{getDoctorById: getDoctorById}
}

func (h *GetDoctorByIdController) GetById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	doctor, err := h.getDoctorById.Execute(int32(idInt))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, doctor)
}