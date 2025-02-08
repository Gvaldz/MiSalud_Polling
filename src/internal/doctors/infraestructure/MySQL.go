package infraestructure

import (
	"database/sql"
	"fmt"
	"log"
	"clinica/src/internal/doctors/domain/entities"
)

type DoctorRepo struct {
	db *sql.DB
}

func NewDoctorRepo(db *sql.DB) *DoctorRepo {
	return &DoctorRepo{db: db}
}

func (r *DoctorRepo) Save(doctor entities.Doctor) error {
	query := "INSERT INTO doctores (cedula, nombres, apellidos, especialidad) VALUES (?, ?, ?, ?)"
	_, err := r.db.Exec(query, doctor.Cedula, doctor.Nombres, doctor.Apellidos, doctor.Especialidad)
	if err != nil {
		return fmt.Errorf("error al guardar doctor: %w", err)
	}
	return nil
}

func (r *DoctorRepo) GetAll() ([]entities.Doctor, error) { 
	query := "SELECT idDoctores, cedula, nombres, apellidos, especialidad FROM doctores"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener doctores: %w", err)
	}
	defer rows.Close()

	var doctors []entities.Doctor
	for rows.Next() {
		var doctor entities.Doctor
		if err := rows.Scan(&doctor.IdDoctores, &doctor.Cedula, &doctor.Nombres, &doctor.Apellidos, &doctor.Especialidad); err != nil {
			return nil, fmt.Errorf("error al escanear doctor: %w", err)
		}
		doctors = append(doctors, doctor)
	}

	return doctors, nil
}

func (r *DoctorRepo) GetById(id int32) (entities.Doctor, error) {
	var doctor entities.Doctor
	query := "SELECT idDoctores, cedula, nombres, apellidos, especialidad FROM doctores WHERE idDoctores = ?"
	err := r.db.QueryRow(query, id).Scan(&doctor.IdDoctores, &doctor.Cedula, &doctor.Nombres, &doctor.Apellidos, &doctor.Especialidad)
	if err != nil {
		return doctor, fmt.Errorf("error al obtener doctor: %w", err)
	}
	return doctor, nil
}

func (r *DoctorRepo) Update(doctor entities.Doctor) error {
	query := "UPDATE doctores SET cedula = ?, nombres = ?, apellidos = ?, especialidad = ? WHERE idDoctores = ?"
	_, err := r.db.Exec(query, doctor.Cedula, doctor.Nombres, doctor.Apellidos, doctor.Especialidad, doctor.IdDoctores)	
	if err != nil {
		return fmt.Errorf("error al actualizar doctor: %w", err)
	}
	log.Println("[DoctorRepo] - Doctor actualizado correctamente")
	return nil
}

func (r *DoctorRepo) Delete(idDoctores int32) error {
	query := "DELETE FROM doctores WHERE idDoctores = ?"
	_, err := r.db.Exec(query, idDoctores)
	if err != nil {
		return fmt.Errorf("error al eliminar doctor: %w", err)
	}
	log.Println("[DoctorRepo] - Doctor eliminado correctamente")
	return nil
}

func (r *DoctorRepo) GetUpdatedSince(since string) ([]entities.Doctor, error) {
	query := "SELECT idDoctores, cedula, nombres, apellidos, especialidad FROM doctores WHERE updated_at > ?"
	rows, err := r.db.Query(query, since)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doctors []entities.Doctor
	for rows.Next() {
		var doctor entities.Doctor
		if err := rows.Scan(&doctor.IdDoctores, &doctor.Cedula, &doctor.Nombres, &doctor.Apellidos, &doctor.Especialidad); err != nil {
			return nil, err
		}
		doctors = append(doctors, doctor)
	}
	return doctors, nil
}