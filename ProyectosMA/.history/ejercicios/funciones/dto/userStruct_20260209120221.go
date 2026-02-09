package dto

type UserStruct struct {
	nombre    string
	apellido1 string
	apellido2 string
	dni       string
	tfno      int
}

func (user *UserStruct) GetNombre() string {
	return user.nombre
}
func (user *UserStruct) SetNombre(nombre string) {
	user.nombre = nombre
}

func (user *UserStruct) GetApellido1() string {
	return user.apellido1
}
func (user *UserStruct) SetApellido1(apellido1 string) {
	user.apellido1 = apellido1
}

func (user UserStruct) GetApellido2() string {
	return user.apellido2
}
func (user *UserStruct) SetApellido2(apellido2 string) {
	user.apellido2 = apellido2
}

func (user UserStruct) GetDni() string {
	return user.dni
}
func (user *UserStruct) SetDni(dni string) {
	user.dni = dni
}

func (user UserStruct) GetTfno() int {
	return user.tfno
}
func (user *UserStruct) SetTfno(tfno int) {
	user.tfno = tfno
}
