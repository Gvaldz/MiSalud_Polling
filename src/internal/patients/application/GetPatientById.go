package application

import (
	"clinica/src/internal/patients/domain"
	"clinica/src/internal/patients/domain/entities"
)

type GetPatientById struct {
	repo domain.IPatient
}

func NewGetPatientById(repo domain.IPatient) *GetPatientById {
	return &GetPatientById{repo: repo}
}

func (g *GetPatientById) Execute(id int32) (entities.Patient, error) {
	return g.repo.GetById(id)
}