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

	persona := &personaStruct{}

	completaDatosStruct(persona)

	getStructFields(persona)

}

func completaDatosStruct(persona *personaStruct) {

	personaValue := reflect.ValueOf(persona).Elem()
	personaType := reflect.TypeOf(persona).Elem()

	for i := 0; i < personaType.NumField(); i++ {
		campo := personaValue.Field(i)
		fieldName := personaType.Field(i).Name

		valorPedido := fmt.Sprintf("Introduzca el valor para %s: ", fieldName)
		fmt.Print(valorPedido)
		var valor string
		fmt.Scanln(&valor)
		campo.SetString(valor)
	}

}

func pideValorParaCampo(campo reflect.Value) string {
	var valor string
	fmt.Printf("Introduzca el valor para %s:", campo.String())
	fmt.Scanln(&valor)
	return valor
}

func getStructFields(p *personaStruct) {
	pv := reflect.ValueOf(p).Elem()
	pt := reflect.TypeOf(p).Elem()
	// Imprime los campos del struct mediante reflection
	println("Datos de la persona:")
	for i := 0; i < pt.NumField(); i++ {
		field := pt.Field(i)
		fmt.Printf("\t%s: %s\n", field.Name, pv.Field(i).String())
	}
}
