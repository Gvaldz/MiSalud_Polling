package controllers

import (
	"clinica/src/internal/patients/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type GetUpdatedPatientsController struct {
	useCase *application.GetUpdatedPatients
}

func NewGetUpdatedPatientsController(useCase *application.GetUpdatedPatients) *GetUpdatedPatientsController {
	return &GetUpdatedPatientsController{useCase: useCase}
}

func (c *GetUpdatedPatientsController) GetUpdates(ctx *gin.Context) {
	since := ctx.Query("since")
	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(2 * time.Second)

	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			ctx.JSON(http.StatusOK, gin.H{"message": "No updates"})
			return
		case <-ticker.C:
			updates, err := c.useCase.Execute(since)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			if len(updates) > 0 {
				ctx.JSON(http.StatusOK, updates)
				return
			}
		}
	}
}
