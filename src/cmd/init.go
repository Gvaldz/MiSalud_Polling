package cmd

import (
	doctorDeps  "clinica/src/internal/doctors/infraestructure/dependences"
	patientDeps "clinica/src/internal/patients/infraestructure/dependences"	
	"clinica/src/core"
	"clinica/src/server"
	"log"
)

func Init() {
	db, err := core.ConnectDB()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}

	doctorDependencies := doctorDeps.NewDoctorDependencies(db)
	doctorRoutes := doctorDependencies.GetRoutes()

	patientDependencies := patientDeps.NewPatientDependencies(db)
	patientRoutes := patientDependencies.GetRoutes()

	server.Run(doctorRoutes, patientRoutes)
}
