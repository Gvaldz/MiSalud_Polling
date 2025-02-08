package application

import (
	"clinica/src/internal/patients/domain"
	"clinica/src/internal/patients/domain/entities"
)

type CreatePatient struct {
	repo domain.IPatient
}

func NewCreatePatient(repo domain.IPatient) *CreatePatient {
	return &CreatePatient{repo: repo}
}

func (cp *CreatePatient) Execute(patient entities.Patient) error{
	
	return cp.repo.Save(patient)
}