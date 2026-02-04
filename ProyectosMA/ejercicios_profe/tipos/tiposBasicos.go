package tipos

import (
	"fmt"

	"math"
)

// Función para mostrar los tipos de datos. No recibe ni devuelve nada.
func MostrarTiposBasicos() {
	/*
		###########
		#VARIABLES#
		###########
	*/
	// Tipos de datos enteros.
	var intNormal int = 58
	var int8bits int8 = -128                  // -128 a 127
	var int16bits int16 = 32767               // -32768 a 32767
	var int32bits int32 = -2147483647         // -2147483648 a 2147483647
	var int64bits int64 = 9223372036854775807 // -9223372036854775808 a 9223372036854775807

	var uintNormal uint = 58
	var uint8bits uint8 = 128                   // 127
	var uint16bits uint16 = 32767               // 32767
	var uint32bits uint32 = 2147483647          // 2147483647
	var uint64bits uint64 = 9223372036854775807 // 9223372036854775807

	// Con la librería [math], puedo acceder a los rangos de los valores.
	var valorMinFloat32 = math.SmallestNonzeroFloat32
	var valorMaxFloat32 = math.MaxFloat32
	var valorMinFloat64 = math.SmallestNonzeroFloat64
	var valorMaxFloat64 = math.MaxFloat64

	// Booleanos.
	var valorVerdadero bool = true
	var valorFalso bool = false

	/*
		#####################
		#IMPRESIÓN VARIABLES#
		#####################
	*/
	fmt.Println("Números enteros + o -")
	fmt.Printf("intNomrmal: %d Typo → %T\n", intNormal, intNormal)
	fmt.Printf("int8bits: %d Typo → %T\n", int8bits, int8bits)
	fmt.Printf("int16bits: %d Typo → %T\n", int16bits, int16bits)
	fmt.Printf("int32bits: %d Typo → %T\n", int32bits, int32bits)
	fmt.Printf("int64bits: %d Typo → %T\n", int64bits, int64bits)
	fmt.Println("Números enteros solo +")
	fmt.Printf("uintNomrmal: %d Typo → %T\n", uintNormal, uintNormal)
	fmt.Printf("uint8bits: %d Typo → %T\n", uint8bits, uint8bits)
	fmt.Printf("uint16bits: %d Typo → %T\n", uint16bits, uint16bits)
	fmt.Printf("uint32bits: %d Typo → %T\n", uint32bits, uint32bits)
	fmt.Printf("uint64bits: %d Typo → %T\n", uint64bits, uint64bits)

	fmt.Printf("El valor mínimo de los float32 es: %d\n", valorMinFloat32)
	fmt.Printf("El valor máximo de los float32 es: %d\n", valorMaxFloat32)
	fmt.Printf("El valor mínimo de los float64 es: %d\n", valorMinFloat64)
	fmt.Printf("El valor mínimo de los float64 es: %d\n", valorMaxFloat64)

	fmt.Println("Tipos booleanos.")

	fmt.Printf("valorVerdadero: [%v] Typo → [%T]\n", valorVerdadero, valorVerdadero)
	fmt.Printf("valorFalso:\t[%v] Typo\t→ [%T]\n", valorFalso, valorFalso)

	/*
		###########
		#TYPO BYTE#
		###########
	*/

	var a byte = 'A'
	fmt.Println(a)

	s := "hola"
	fmt.Println(s[0])

	var r rune = '☻'
	fmt.Println(r)
}
