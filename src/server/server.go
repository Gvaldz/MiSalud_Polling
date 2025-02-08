package server

import (
	"github.com/gin-gonic/gin"
	doctorRouters "clinica/src/internal/doctors/infraestructure/routers"
	patientRouters "clinica/src/internal/patients/infraestructure/routers"
)

func Run(
	doctorRoutes *doctorRouters.DoctorRoutes,
	patientRoutes *patientRouters.PatientRoutes,
) {
	r := gin.Default()

	doctorRoutes.AttachRoutes(r)
	patientRoutes.AttachRoutes(r)

	r.Run(":8080")
}
