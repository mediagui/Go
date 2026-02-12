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

import (
	"fmt"
	"reflect"
)

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

	completaDatosStruct(persona)

	getStructFields(persona)

}

func completaDatosStruct(persona personaStruct) {

	personaValue := reflect.ValueOf(persona)

	for _, campo := range personaValue.Fields() {

		// campo := campo

		valorPedido := pideValorParaCampo(campo)

		campo.SetString(valorPedido)
	}

}

func pideValorParaCampo(campo reflect.Value) string {
	var valor string
	fmt.Printf("Introduzca el valor para %s:", campo.String())
	fmt.Scanln(&valor)
	return valor
}

func getStructFields(p personaStruct) {
	pv := reflect.ValueOf(p)
	pt := reflect.TypeOf(p)
	// Imprime los campos del struct mediante reflection
	println("Datos de la persona:")
	for field := range pt.Fields() {
		fmt.Printf("\t%s: %s\n", field.Name, pv.FieldByName(field.Name).String())
	}
}
