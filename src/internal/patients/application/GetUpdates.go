package application

import (
	"clinica/src/internal/patients/domain"
	"clinica/src/internal/patients/domain/entities"
)

type GetUpdatedPatients struct {
	patientRepo domain.IPatient
}

func NewGetUpdatedPatients(patientRepo domain.IPatient) *GetUpdatedPatients {
	return &GetUpdatedPatients{patientRepo: patientRepo}
}

func (u *GetUpdatedPatients) Execute(since string) ([]entities.Patient, error) {
	return u.patientRepo.GetUpdatedPatients(since)
}
