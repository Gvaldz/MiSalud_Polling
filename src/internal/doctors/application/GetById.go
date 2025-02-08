package application

import (
	"clinica/src/internal/doctors/domain"
	"clinica/src/internal/doctors/domain/entities"
)

type GetDoctorById struct {
	repo domain.IDoctor
}

func NewGetDoctorById(repo domain.IDoctor) *GetDoctorById {
	return &GetDoctorById{repo: repo}
}

func (g *GetDoctorById) Execute(id int32) (entities.Doctor, error) {
	return g.repo.GetById(id)
}