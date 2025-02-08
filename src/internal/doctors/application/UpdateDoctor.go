package application

import (
	"clinica/src/internal/doctors/domain"
	"clinica/src/internal/doctors/domain/entities"
)

type UpdateDoctor struct {
	db domain.IDoctor
}

func NewUpdateDoctor(db domain.IDoctor) *UpdateDoctor {
	return &UpdateDoctor{db: db}
}

func (lp *UpdateDoctor) Execute(doctor entities.Doctor) error{
	return lp.db.Update(doctor)
}