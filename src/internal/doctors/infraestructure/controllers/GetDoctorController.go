package controllers

import (
	"clinica/src/internal/doctors/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllDoctorController struct {
	listDoctors  *application.ListDoctors
}

func NewGetAllDoctorController(list *application.ListDoctors) *GetAllDoctorController {
	return &GetAllDoctorController{
		listDoctors:  list,
	}
}

func (h *GetAllDoctorController) GetAll(c *gin.Context) {
	doctors, err := h.listDoctors.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, doctors)
}