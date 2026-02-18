package dto

type userStruct struct {
	nombre    string
	apellido1 string
	apellido2 string
	dni       string
	tfno      int
}

func (user *userStruct) GetNombre() string {
	return user.nombre
}

func (user *userStruct) SetNombre(nombre string) {
	user.nombre = nombre
}

func (user userStruct) GetApellido1() string {
	return user.apellido1
}

func (user userStruct) GetApellido2() string {
	return user.apellido2
}

func (user userStruct) GetDni() string {
	return user.dni
}

func (user userStruct) GetTfno() int {
	return user.tfno
}
