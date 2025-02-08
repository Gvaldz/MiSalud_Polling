package controllers

import (
	"clinica/src/internal/doctors/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteDoctorController struct {
	deleteDoctor *application.DeleteDoctor
}

func NewDeleteDoctorController(delete *application.DeleteDoctor) *DeleteDoctorController {
	return &DeleteDoctorController{
		deleteDoctor: delete,
	}
}

func (h *DeleteDoctorController) Delete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.deleteDoctor.Execute(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Doctor eliminado correctamente"})
}