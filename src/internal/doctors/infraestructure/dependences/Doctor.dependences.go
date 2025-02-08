package dependences

import (
	"clinica/src/internal/doctors/application"
	"clinica/src/internal/doctors/infraestructure"
	"clinica/src/internal/doctors/infraestructure/controllers"
	"clinica/src/internal/doctors/infraestructure/routers"
	"database/sql"
)

type DoctorDependencies struct {
	DB *sql.DB
}

func NewDoctorDependencies(db *sql.DB) *DoctorDependencies {
	return &DoctorDependencies{DB: db}
}

func (d *DoctorDependencies) GetRoutes() *routers.DoctorRoutes {
	doctorRepo := infraestructure.NewDoctorRepo(d.DB)

	createDoctorUseCase := application.NewCreateDoctor(doctorRepo)
	listDoctorsUseCase := application.NewListDoctors(doctorRepo)
	getDoctorByIdUseCase := application.NewGetDoctorById(doctorRepo)
	updateDoctorUseCase := application.NewUpdateDoctor(doctorRepo)
	deleteDoctorUseCase := application.NewDeleteDoctor(doctorRepo)
	getUpdatedDoctorsUseCase := application.NewGetUpdatedDoctors(doctorRepo)

	createDoctorController := controllers.NewCreateDoctorController(createDoctorUseCase)
	listDoctorsController := controllers.NewGetAllDoctorController(listDoctorsUseCase)
	getDoctorController := controllers.NewGetDoctorByIdController(getDoctorByIdUseCase)
	updateDoctorController := controllers.NewUpdateDoctorController(updateDoctorUseCase)
	deleteDoctorController := controllers.NewDeleteDoctorController(deleteDoctorUseCase)
	getUpdatedDoctorController := controllers.NewGetUpdatedDoctorsController(getUpdatedDoctorsUseCase)

	return routers.NewDoctorRoutes(
		createDoctorController,
		listDoctorsController,
		getDoctorController,
		updateDoctorController,
		deleteDoctorController,
		getUpdatedDoctorController,
	)
}