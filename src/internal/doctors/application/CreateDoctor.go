package application

import (
	"clinica/src/internal/doctors/domain"
	"clinica/src/internal/doctors/domain/entities"
)

type CreateDoctor struct {
	repo domain.IDoctor
}

func NewCreateDoctor(repo domain.IDoctor) *CreateDoctor {
	return &CreateDoctor{repo: repo}
}

func (cp *CreateDoctor) Execute(doctor entities.Doctor) error{
	
	return cp.repo.Save(doctor)
}