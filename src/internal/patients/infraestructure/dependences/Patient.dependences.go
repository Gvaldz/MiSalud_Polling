package dependences

import (
	"clinica/src/internal/patients/application"
	"clinica/src/internal/patients/infraestructure"
	"clinica/src/internal/patients/infraestructure/controllers"
	"clinica/src/internal/patients/infraestructure/routers"
	"database/sql"
)

type PatientDependencies struct {
	DB *sql.DB
}

func NewPatientDependencies(db *sql.DB) *PatientDependencies {
	return &PatientDependencies{DB: db}
}

func (d *PatientDependencies) GetRoutes() *routers.PatientRoutes {
	patientRepo := infraestructure.NewPatientRepo(d.DB)

	createPatientUseCase := application.NewCreatePatient(patientRepo)
	listPatientsUseCase := application.NewListPatients(patientRepo)
	getPatientByIdUseCase := application.NewGetPatientById(patientRepo)
	updatePatientUseCase := application.NewUpdatePatients(patientRepo)
	deletePatientUseCase := application.NewDeletePatient(patientRepo)
	getUpdatesUseCase := application.NewGetUpdatedPatients(patientRepo)
	getCountUseCase := application.NewGetPatientCount(patientRepo)

	createPatientController := controllers.NewCreatePatientController(createPatientUseCase)
	listPatientsController := controllers.NewGetAllPatientsController(listPatientsUseCase)
	getPatientController := controllers.NewGetPatientByIdController(getPatientByIdUseCase)
	updatePatientController := controllers.NewUpdatePatientsController(updatePatientUseCase)
	deletePatientController := controllers.NewDeletePatientController(deletePatientUseCase)
	getUpdatesController := controllers.NewGetUpdatedPatientsController(getUpdatesUseCase)
	getCountController := controllers.NewGetPatientCountController(getCountUseCase)

	return routers.NewPatientRoutes(
		createPatientController,
		listPatientsController,
		getPatientController,
		updatePatientController,
		deletePatientController,
		getUpdatesController,
		getCountController,
	)
}