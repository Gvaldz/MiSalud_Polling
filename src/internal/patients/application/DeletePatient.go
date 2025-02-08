package application

import (
	"clinica/src/internal/patients/domain"
)

type DeletePatient struct {
	repo domain.IPatient
}

func NewDeletePatient(repo domain.IPatient) *DeletePatient {
	return &DeletePatient{repo: repo}
}

func (lp *DeletePatient) Execute(id int32) error{
	return lp.repo.Delete(id)
}