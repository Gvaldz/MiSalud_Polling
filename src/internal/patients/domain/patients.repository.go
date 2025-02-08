package domain

import "clinica/src/internal/patients/domain/entities"

type IPatient interface {
	GetAll() ([]entities.Patient, error)
	GetById(id int32) (entities.Patient, error)
	Save(patient entities.Patient) error
	Update(patient entities.Patient) error
	Delete(idPatients int32) error 
	GetUpdatedPatients(since string) ([]entities.Patient, error)
	Count() (int, error)
}