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
	Nombre          string `etiqueta:"Nombre: "`
	Apellidos       string `etiqueta:"Apellidos: "`
	DNI             string `etiqueta:"N.I.F.: "`
	Edad            string `etiqueta:"Añitos: "`
	EstadoCivil     string `etiqueta:"E. Civil: "`
	FechaNacimiento string `etiqueta:"F. Nacimiento: "`
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
		nombreCampo := personaType.Field(i).Name

		valor := pideValorParaCampo(nombreCampo)

		campo.SetString(valor)
	}

}

func pideValorParaCampo(fieldName string) string {
	fmt.Printf("Introduzca el valor para %s: ", fieldName)
	var valor string
	fmt.Scanln(&valor)
	return valor
}

func getStructFields(p *personaStruct) {
	pv := reflect.ValueOf(p).Elem()
	pt := reflect.TypeFor[personaStruct]()

	// Imprime los campos del struct mediante reflection
	println("Datos de la persona:")
	for i := 0; i < pt.NumField(); i++ {
		field := pt.Field(i)
		fmt.Printf("\t%s\t%s\n", field.Tag.Get("etiqueta"), pv.Field(i).String())
	}
}
