package application

import "clinica/src/internal/patients/domain"

type GetPatientCount struct {
	patientRepo domain.IPatient
}

func NewGetPatientCount(patientRepo domain.IPatient) *GetPatientCount {
	return &GetPatientCount{patientRepo: patientRepo}
}

func (u *GetPatientCount) Execute() (int, error) {
	return u.patientRepo.Count()
}
