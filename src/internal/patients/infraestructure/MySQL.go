package infraestructure

import (
	"database/sql"
	"fmt"
	"log"
	"clinica/src/internal/patients/domain/entities"
)

type PatientRepo struct {
	db *sql.DB
}

func NewPatientRepo(db *sql.DB) *PatientRepo {
	return &PatientRepo{db: db}
}

func (r *PatientRepo) Save(patient entities.Patient) error {
	query := "INSERT INTO pacientes (nombres, apellidos, nacimiento, genero, peso, estatura) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := r.db.Exec(query, patient.Nombres, patient.Apellidos, patient.Nacimiento, patient.Genero, patient.Peso, patient.Estatura)
	if err != nil {
		return fmt.Errorf("error al guardar paciente: %w", err)
	}
	
	return nil
}

func (r *PatientRepo) GetAll() ([]entities.Patient, error) { 
	query := "SELECT idPatients, nombres, apellidos, nacimiento, genero, peso, estatura FROM pacientes"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener pacientes: %w", err)
	}
	defer rows.Close()

	var patients []entities.Patient
	for rows.Next() {
		var patient entities.Patient
		if err := rows.Scan(&patient.IdPatients, &patient.Nombres, &patient.Apellidos, &patient.Nacimiento, &patient.Genero, &patient.Peso, &patient.Estatura); err != nil {
			return nil, fmt.Errorf("error al escanear paciente: %w", err)
		}
		patients = append(patients, patient)
	}

	return patients, nil
}

func (r *PatientRepo) GetById(id int32) (entities.Patient, error) {
	var patient entities.Patient
	query := "SELECT idPatients, nombres, apellidos, nacimiento, genero, peso, estatura FROM pacientes WHERE idPatients = ?"
	err := r.db.QueryRow(query, id).Scan(&patient.IdPatients, &patient.Nombres, &patient.Apellidos, &patient.Nacimiento, &patient.Genero, &patient.Peso, &patient.Estatura)
	if err != nil {
		return patient, fmt.Errorf("error al obtener paciente: %w", err)
	}
	return patient, nil
}

func (r *PatientRepo) Update(patient entities.Patient) error {
	query := "UPDATE pacientes SET nombres = ?, apellidos = ?, nacimiento = ?, genero = ?, peso = ?, estatura = ? WHERE idPatients = ?"
	_, err := r.db.Exec(query, patient.Nombres, patient.Apellidos, patient.Nacimiento, patient.Genero, patient.Peso, patient.Estatura, patient.IdPatients)	
	if err != nil {
		return fmt.Errorf("error al actualizar paciente: %w", err)
	}
	log.Println("[PatientRepo] - Paciente actualizado correctamente")
	return nil
}

func (r *PatientRepo) Delete(idPatients int32) error {
	query := "DELETE FROM pacientes WHERE idPatients = ?"
	_, err := r.db.Exec(query, idPatients)
	if err != nil {
		return fmt.Errorf("error al eliminar paciente: %w", err)
	}
	log.Println("[PatientRepo] - Paciente eliminado correctamente")
	return nil
}

func (r *PatientRepo) GetUpdatedPatients(since string) ([]entities.Patient, error) {
	query := "SELECT idpatients, nombres, apellidos, nacimiento, genero, peso, estatura FROM pacientes WHERE updated_at > ?"
	rows, err := r.db.Query(query, since)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []entities.Patient
	for rows.Next() {
		var p entities.Patient
		if err := rows.Scan(&p.IdPatients, &p.Nombres, &p.Apellidos, &p.Nacimiento, &p.Genero, &p.Peso, &p.Estatura); err != nil {
			return nil, err
		}
		patients = append(patients, p)
	}

	return patients, nil
}

func (r *PatientRepo) Count() (int, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM pacientes").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
