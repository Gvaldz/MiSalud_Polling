package application

import (
	"clinica/src/internal/patients/domain"
	"clinica/src/internal/patients/domain/entities"
)

type UpdatePatients struct {
	db domain.IPatient
}

func NewUpdatePatients(db domain.IPatient) *UpdatePatients {
	return &UpdatePatients{db: db}
}

func (lp *UpdatePatients) Execute(patients entities.Patient) error{
	return lp.db.Update(patients)
}