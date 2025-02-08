package application

import (
	"clinica/src/internal/doctors/domain"
	"clinica/src/internal/doctors/domain/entities"

)

type GetUpdatedDoctors struct {
	repo domain.IDoctor
}

func NewGetUpdatedDoctors(repo domain.IDoctor) *GetUpdatedDoctors {
	return &GetUpdatedDoctors{repo: repo}
}

func (u *GetUpdatedDoctors) Execute(since string) ([]entities.Doctor, error) {
	return u.repo.GetUpdatedSince(since)
}
