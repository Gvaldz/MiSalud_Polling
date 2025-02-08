package entities

type Doctor struct {
	IdDoctores int32
	Cedula int32
	Nombres string
	Apellidos string
	Especialidad string
}

func NewDoctor(idDoctores int32, cedula int32, nombres string, apellidos string, especialidad string) *Doctor{
	return &Doctor{IdDoctores: idDoctores, Cedula: cedula, Nombres: nombres, Apellidos: apellidos, Especialidad: especialidad}
}

