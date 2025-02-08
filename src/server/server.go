package server

import (
	doctorRouters "clinica/src/internal/doctors/infraestructure/routers"
	patientRouters "clinica/src/internal/patients/infraestructure/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(
	doctorRoutes *doctorRouters.DoctorRoutes,
	patientRoutes *patientRouters.PatientRoutes,
) {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	
	}))

	doctorRoutes.AttachRoutes(r)
	patientRoutes.AttachRoutes(r)

	r.Run(":8080")
}
