// Crear una aplicación  en go, que permita crear objetos de tipos [Persona]. Después se mostrará cada uno de los objetos creados.
// Los atributos para cada persona, serán:

// v  Nombre.
// v  Apellidos.
// v  DNI.
// v  Edad.
// v  Estado civil.
// v  Año de nacimiento.

// Por cada Persona creada, se deberán solicitar los datos pertinentes.

// Paquetes usados:
//   - fmt: funciones para entrada/salida y formato. Ver: https://pkg.go.dev/fmt
//   - reflect: inspección y modificación de tipos y valores en tiempo de ejecución.
//     Ver: https://pkg.go.dev/reflect
package main

import (
	"fmt"
	"reflect"

	"github.com/mediagui/structs/util"
)

// `personaStruct` define los campos que representan los datos de una persona.
// Cada campo lleva una `struct tag` con la clave "etiqueta" usada para mostrar
// un texto descriptivo cuando se imprimen los valores.
// Más sobre struct tags: https://pkg.go.dev/reflect#StructTag.Get
type personaStruct struct {
	Nombre          string `etiqueta:"Nombre: "`
	Apellidos       string `etiqueta:"Apellidos: "`
	DNI             string `etiqueta:"N.I.F.: "`
	Edad            string `etiqueta:"Añitos: "`
	EstadoCivil     string `etiqueta:"E. Civil: "`
	FechaNacimiento string `etiqueta:"F. Nacimiento: "`
}

func main() {

	// Crea un puntero a `personaStruct` y pasa ese puntero a las funciones
	// que rellenan los datos y muestran los campos.

	persona := &personaStruct{}

	// Pasamos un puntero porque las funciones usan reflection y requieren
	// valores direccionables ( `reflect.Value.Elem()` ).
	// Ver: https://pkg.go.dev/reflect#Value.Elem

	completaDatosStruct(persona)
	getStructFields(persona)

}

// completaDatosStruct rellena los campos de `persona` leyendo valores desde
// la entrada estándar. Usa reflection para obtener un `reflect.Value`
// direccionable sobre el struct (mediante `reflect.ValueOf(...).Elem()`),
// y un `reflect.Type` para iterar los campos (`NumField`, `Field`).
// Por cada campo solicita el valor al usuario y lo asigna con
//
//	`reflect.Value.SetString` (requiere que el campo sea asignable y de tipo string).
//
// Enlaces:
//   - reflect.ValueOf: https://pkg.go.dev/reflect#ValueOf
//   - reflect.TypeOf:  https://pkg.go.dev/reflect#TypeOf
//   - Value.Elem/Value.Field/Value.SetString: https://pkg.go.dev/reflect#Value
//   - Type.NumField/Type.Field: https://pkg.go.dev/reflect#Type
func completaDatosStruct(persona *personaStruct) {

	// Obtener el `reflect.Value` direccionable del puntero `persona`.
	//  `reflect.ValueOf(persona)` devuelve un Value que contiene el puntero;
	//  `Elem()` obtiene el valor al que apunta (el struct) para permitir modificaciones.
	// Ver: https://pkg.go.dev/reflect#ValueOf
	personaValue := reflect.ValueOf(persona).Elem()

	// Obtener el `reflect.Type` del struct apuntado por `persona`.
	// `TypeOf(...).Elem()` se usa cuando se parte de un puntero a un struct.
	// Ver: https://pkg.go.dev/reflect#TypeOf
	personaType := reflect.TypeOf(persona).Elem()

	// Iteramos cada campo del tipo struct usando NumField/Field.
	// Ver: https://pkg.go.dev/reflect#Type.NumField
	for i := 0; i < personaType.NumField(); i++ {
		// `personaValue.Field(i)` devuelve el i-ésimo campo como reflect.Value
		// (permite leer o escribir el campo si es direccionable y asignable).
		// Ver: https://pkg.go.dev/reflect#Value.Field
		campo := personaValue.Field(i)

		// `personaType.Field(i).Name` obtiene el nombre declarado del campo.
		nombreCampo := personaType.Field(i).Name

		// Pedimos al usuario el valor para el campo concreto.
		valor := pideValorParaCampo(nombreCampo)

		// Asignamos el string al campo usando SetString (requiere que el
		// campo sea de tipo string y asignable).
		// Ver: https://pkg.go.dev/reflect#Value.SetString
		campo.SetString(valor)
	}

}

// pideValorParaCampo muestra un prompt con el nombre del campo y lee la
// entrada del usuario desde la entrada estándar, devolviendo la cadena
// leída.
//
// Detalles:
// - Usa `fmt.Printf` para mostrar el prompt (ver https://pkg.go.dev/fmt#Printf)
// - Usa `fmt.Scanln` para leer la entrada hasta newline (ver https://pkg.go.dev/fmt#Scanln)
//
// Nota: `fmt.Scanln` divide la entrada por espacios; si necesita leer líneas
// completas con espacios considere usar `bufio.NewReader(os.Stdin).ReadString('\n')`.
func pideValorParaCampo(fieldName string) string {

	// Muestra un prompt con `fmt.Printf` y lee la entrada con `fmt.Scanln`.
	// `fmt.Printf` — formatea y escribe en la salida estándar.
	// Ver: https://pkg.go.dev/fmt#Printf
	fmt.Printf("Introduzca el valor para %s: ", fieldName)

	var valor string

	// `fmt.Scanln` lee de la entrada estándar separando por espacios hasta newline.
	// Nota: si se necesitan líneas completas con espacios, usar `bufio.Reader`.
	// Ver: https://pkg.go.dev/fmt#Scanln
	fmt.Scanln(&valor)

	return valor
}

// getStructFields muestra por consola las etiquetas y los valores de cada
// campo del struct `personaStruct`. Obtiene el `reflect.Type` del tipo para
// leer las `StructTag` (ej.
//
//	`field.Tag.Get("etiqueta")`
//
// ) y usa el
// `reflect.Value` direccionable para leer los valores actuales
// ( `Value.Field(i).String()` ).
//
// Enlaces:
//   - reflect.TypeFor / TypeOf: https://pkg.go.dev/reflect#TypeFor
//   - StructTag.Get:        https://pkg.go.dev/reflect#StructTag.Get
//   - Value.Field/Value.String: https://pkg.go.dev/reflect#Value
func getStructFields(p *personaStruct) {

	// Valor direccionable del struct apuntado por `p`.
	// Ver: https://pkg.go.dev/reflect#ValueOf
	pv := reflect.ValueOf(p).Elem()

	// Obtener el reflect.Type para el tipo `personaStruct`.
	// `reflect.TypeFor[T]()` es una forma genérica de obtener el Type.
	// Alternativas: `reflect.TypeOf(personaStruct{})` o
	// `reflect.TypeOf((*personaStruct)(nil)).Elem()`.
	// Ver: https://pkg.go.dev/reflect#TypeFor
	pt := reflect.TypeFor[personaStruct]()

	// Imprime los campos del struct mediante reflection
	println("Datos de la persona:")
	for i := 0; i < pt.NumField(); i++ {
		// `pt.Field(i)` devuelve información del campo (nombre, tipo, tags).
		// Ver: https://pkg.go.dev/reflect#Type.Field
		field := pt.Field(i)

		// `field.Tag.Get("etiqueta")` obtiene la tag con clave "etiqueta".
		// `pv.Field(i).String()` obtiene el valor actual del campo como string.
		// Ver: https://pkg.go.dev/reflect#StructTag.Get
		// Ver: https://pkg.go.dev/reflect#Value.String

		fmt.Printf("\t%s\t%s\n", field.Tag.Get("etiqueta"), pv.Field(i).String())

		if field.Name == "DNI" {
			fmt.Printf("\t - NIF correcto: %v\n", util.ValidateNIF(pv.Field(i).String()))
		}

	}
}
