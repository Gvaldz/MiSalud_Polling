package application

import (
	"clinica/src/internal/patients/domain"
	"clinica/src/internal/patients/domain/entities"
)

type ListPatients struct {
	repo domain.IPatient
}

func NewListPatients(repo domain.IPatient) *ListPatients {
	return &ListPatients{repo: repo}
}

func (lp *ListPatients) Execute() ([]entities.Patient, error) {
    return lp.repo.GetAll()
}
