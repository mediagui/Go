// Crear una aplicación  en go, que permita crear objetos de tipos [Persona]. Después se mostrará cada uno de los objetos creados.
// Los atributos para cada persona, serán:

// v  Nombre.
// v  Apellidos.
// v  DNI.
// v  Edad.
// v  Estado civil.
// v  Año de nacimiento.

// Por cada Persona creada, se deberán solicitar los datos pertinentes.

package main

type personaStruct struct {
	Nombre          string
	Apellidos       string
	DNI             string
	Edad            string
	EstadoCivil     string
	FechaNacimiento string
}

func main() {

	var persona personaStruct

}
