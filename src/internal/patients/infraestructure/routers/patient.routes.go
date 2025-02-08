package routers

import (
	"github.com/gin-gonic/gin"
	"clinica/src/internal/patients/infraestructure/controllers"
)

type PatientRoutes struct {
	CreatePatientController *controllers.CreatePatientController
	GetAllPatientsController *controllers.GetAllPatientsController
	GetPatientByIdController *controllers.GetPatientByIdController
	UpdatePatientController *controllers.UpdatePatientsController
	DeletePatientController *controllers.DeletePatientController
	GetUpdatedPatientsController *controllers.GetUpdatedPatientsController
	GetCountController *controllers.GetPatientCountController
}

func NewPatientRoutes(
	createPatientController *controllers.CreatePatientController,
	getAllPatientsController *controllers.GetAllPatientsController,
	getPatientByIdController *controllers.GetPatientByIdController,
	updatePatientController *controllers.UpdatePatientsController,
	deletePatientController *controllers.DeletePatientController,
	getUpdatedPatientsController *controllers.GetUpdatedPatientsController,
	getCountController * controllers.GetPatientCountController,

) *PatientRoutes {
	return &PatientRoutes{
		CreatePatientController: createPatientController,
		GetPatientByIdController: getPatientByIdController,
		GetAllPatientsController: getAllPatientsController,
		UpdatePatientController: updatePatientController,
		DeletePatientController: deletePatientController,
		GetUpdatedPatientsController: getUpdatedPatientsController,
		GetCountController: getCountController,
	}
}

func (r *PatientRoutes) AttachRoutes(router *gin.Engine) {
	PatientsGroup := router.Group("/patients")
	{
		PatientsGroup.POST("", r.CreatePatientController.Create)
		PatientsGroup.GET("", r.GetAllPatientsController.GetAll)
		PatientsGroup.GET("/:id", r.GetPatientByIdController.GetById)	
		PatientsGroup.PUT("/:id", r.UpdatePatientController.Update)
		PatientsGroup.DELETE("/:id", r.DeletePatientController.Delete)
		PatientsGroup.GET("/updates", r.GetUpdatedPatientsController.GetUpdates) 
		PatientsGroup.GET("/count", r.GetCountController.GetCount) 
	}
}