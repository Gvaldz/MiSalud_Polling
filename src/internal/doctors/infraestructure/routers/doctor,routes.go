package routers

import (
	"github.com/gin-gonic/gin"
	"clinica/src/internal/doctors/infraestructure/controllers"
)

type DoctorRoutes struct {
	CreateDoctorController *controllers.CreateDoctorController
	GetAllDoctorsController *controllers.GetAllDoctorController
	GetDoctorByIdController *controllers.GetDoctorByIdController
	UpdateDoctorController *controllers.UpdateDoctorController
	DeleteDoctorController *controllers.DeleteDoctorController
	GetUpdatedDoctorsController *controllers.GetUpdatedDoctorsController
}

func NewDoctorRoutes(
	createDoctorController *controllers.CreateDoctorController,
	getAllDoctorsController *controllers.GetAllDoctorController,
	getDoctorByIdController *controllers.GetDoctorByIdController,
	updateDoctorController *controllers.UpdateDoctorController,
	deleteDoctorController *controllers.DeleteDoctorController,
	getUpdatedDoctorsController *controllers.GetUpdatedDoctorsController,
) *DoctorRoutes {

	return &DoctorRoutes{
		CreateDoctorController: createDoctorController,
		GetAllDoctorsController: getAllDoctorsController,
		GetDoctorByIdController: getDoctorByIdController,
		UpdateDoctorController: updateDoctorController,
		DeleteDoctorController: deleteDoctorController,
		GetUpdatedDoctorsController: getUpdatedDoctorsController,
	}
}

func (r *DoctorRoutes) AttachRoutes(router *gin.Engine) {
	DoctorsGroup := router.Group("/doctors")
	{
		DoctorsGroup.POST("", r.CreateDoctorController.Create)
		DoctorsGroup.GET("", r.GetAllDoctorsController.GetAll)
		DoctorsGroup.GET("/:id", r.GetDoctorByIdController.GetById)
		DoctorsGroup.PUT("/:id", r.UpdateDoctorController.Update)
		DoctorsGroup.DELETE("/:id", r.DeleteDoctorController.Delete)
		DoctorsGroup.GET("/updates", r.GetUpdatedDoctorsController.GetUpdatedDoctors)
	}
}