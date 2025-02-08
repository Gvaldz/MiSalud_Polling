package application

import (
	"clinica/src/internal/patients/domain"
	"clinica/src/internal/patients/domain/entities"
)

type GetPatientStatusUseCase struct {
	patientRepo domain.IPatient
}

func NewGetPatientStatusUseCase(patientRepo domain.IPatient) *GetPatientStatusUseCase {
	return &GetPatientStatusUseCase{patientRepo: patientRepo}
}

func (uc *GetPatientStatusUseCase) Execute(id int32) (entities.Patient, error) {
	return uc.patientRepo.GetById(id)
}