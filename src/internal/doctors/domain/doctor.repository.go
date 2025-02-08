package domain

import "clinica/src/internal/doctors/domain/entities"

type IDoctor interface {
	GetAll() ([]entities.Doctor, error)
	GetById(id int32) (entities.Doctor, error)
	Save(doctor entities.Doctor) error
	Update(doctor entities.Doctor) error
	Delete(idDoctores int32) error
	GetUpdatedSince(since string) ([]entities.Doctor, error) 
}