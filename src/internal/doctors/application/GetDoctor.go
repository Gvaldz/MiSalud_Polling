package application

import (
	"clinica/src/internal/doctors/domain"
	"clinica/src/internal/doctors/domain/entities"
)

type ListDoctors struct {
	repo domain.IDoctor
}

func NewListDoctors(repo domain.IDoctor) *ListDoctors {
	return &ListDoctors{repo: repo}
}

func (lp *ListDoctors) Execute() ([]entities.Doctor, error) {
    return lp.repo.GetAll()
}
