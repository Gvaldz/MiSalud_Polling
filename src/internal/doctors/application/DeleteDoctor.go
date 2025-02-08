package application

import (
	"clinica/src/internal/doctors/domain"
)

type DeleteDoctor struct {
	repo domain.IDoctor
}

func NewDeleteDoctor(repo domain.IDoctor) *DeleteDoctor {
	return &DeleteDoctor{repo: repo}
}

func (lp *DeleteDoctor) Execute(id int32) error{
	return lp.repo.Delete(id)
}