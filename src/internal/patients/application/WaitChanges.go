package application

import (
	"time"
	"clinica/src/internal/patients/domain"
	"clinica/src/internal/patients/domain/entities"
)

type WaitForPatientChangeUseCase struct {
	patientRepo domain.IPatient
}

func NewWaitForPatientChangeUseCase(patientRepo domain.IPatient) *WaitForPatientChangeUseCase {
	return &WaitForPatientChangeUseCase{patientRepo: patientRepo}
}

func (uc *WaitForPatientChangeUseCase) Execute(id int32, timeout time.Duration) (entities.Patient, error) {
	startTime := time.Now()
	for {
		patient, err := uc.patientRepo.GetById(id)
		if err != nil {
			return entities.Patient{}, err
		}

		if time.Since(startTime) > timeout {
			return patient, nil
		}

		time.Sleep(1 * time.Second)
	}
}