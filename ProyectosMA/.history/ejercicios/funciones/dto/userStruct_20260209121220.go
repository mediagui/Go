package dto

type UsuarioStruct struct {
	nombre    string
	apellido1 string
	apellido2 string
	dni       string
	tfno      int
}

func (user *UsuarioStruct) GetNombre() string {
	return user.nombre
}
func (user *UsuarioStruct) SetNombre(nombre string) {
	user.nombre = nombre
}

func (user *UsuarioStruct) GetApellido1() string {
	return user.apellido1
}
func (user *UsuarioStruct) SetApellido1(apellido1 string) {
	user.apellido1 = apellido1
}

func (user UsuarioStruct) GetApellido2() string {
	return user.apellido2
}
func (user *UsuarioStruct) SetApellido2(apellido2 string) {
	user.apellido2 = apellido2
}

func (user UsuarioStruct) GetDni() string {
	return user.dni
}
func (user *UsuarioStruct) SetDni(dni string) {
	user.dni = dni
}

func (user UsuarioStruct) GetTfno() int {
	return user.tfno
}
func (user *UsuarioStruct) SetTfno(tfno int) {
	user.tfno = tfno
}
