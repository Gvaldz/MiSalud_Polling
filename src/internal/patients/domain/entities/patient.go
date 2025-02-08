package entities

type Patient struct {
	IdPatients int32
	Nombres string
	Apellidos string
	Nacimiento string
	Genero string
	Peso float32
	Estatura float32
}

func NewPatient(idPatients int32, nombres string, apellidos string, nacimiento string, genero string, peso float32, estatura float32) *Patient{
	return &Patient{IdPatients: idPatients, Nombres: nombres, Apellidos: apellidos, Nacimiento: nacimiento, Genero: genero, Peso: peso, Estatura: estatura}
}