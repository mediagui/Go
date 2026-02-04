// Indico el paquete al que pertence el archivo.
package variables

// Paquetes externos.
import "fmt"

/*
Tenemos distintos tipos de datos con los que podemos trabajar en Go y estos a su vez se deglosan en otros ditintos.

→ Cadenas de textos (String)
→ Números enteros (int | integer)
	→ uint8:	8-bits(0 a 255)
	→ uint16:	16-bits(0 a 65535)
	→ uint32:	32-bits(0 a 4294967295)
	→ uint64:	64-bits(0 a 18446744073709551615)
	→ int8:		8-bits (-128 a 127)
	→ int16:	16-bits (-32768 a 32767)
	→ int32:	32-bits (-2147483648 a 2147483647)
	→ int64:	64-bits (-9223372036854775808 a 9223372036854775807)
→ Números flotantes (float)
	→ float32
	→ float64
	→ complex64
	→ complex128
→ Booleanos (boolean)
	→ true
	→ false
*/
// Función [MostrarDatos()] que imprimirá cada datos epecificado. Esta función no recibe ningún parámetro.
func MostrarDatos() {
	fmt.Println("Soy un dato de tipo String") // String.
	// Números
	fmt.Println(25)  // int
	fmt.Println(2.5) // float64 o float32
	// Booleanos
	fmt.Println(true)  // Booleano
	fmt.Println(false) // Booleano
}
