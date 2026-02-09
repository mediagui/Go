// Pauqete al que pertenece el programa.
package arrays

// Importaciones.
import "fmt"

/*
Los arrays son un contenedor, que contiene otros contenedores. Básicamente, una estructura de datos.

Es como un armario con cajones, el armario en sí es el array y cada cajón, es una "variable".
*/

func TratarArrays() {
	/*
		#MANERA 1#

		Definición creando un array vacío
	*/
	// Declaro un array vacío.
	var arrayUno [2]string

	// Añado datos al array.
	arrayUno[0] = "dato 01" // Indice 0 Posición 1
	arrayUno[1] = "dato 02" // Indice 0 Posición 2

	// Imprimo el contenido del [arrayUno].
	fmt.Println(arrayUno)

	/*
		#MANERA 2#

		Definición creando un array con valores iniciales.
	*/
	arrayDos := [3]int{1, 2, 3}

	// Imprimo el contenido del [arrayDos].
	fmt.Println(arrayDos)

	// Modificar datos.
	arrayDos[2] = 300

	// Imprimo el contenido del [arrayDos].
	fmt.Println(arrayDos)

	// Imprimo cada número por separado.
	fmt.Println(arrayDos[0])
	fmt.Println(arrayDos[1])
	fmt.Println(arrayDos[2])

	// Inicializo un array, sin tener que indicar el tamaño.
	frutas := [...]string{"Manzana", "Tomate", "Piña", "Kiwi"}

	// Imprimo el array de frutas[]
	fmt.Println(frutas)

	// Obtengo la longitud del array, con la función [len()] y la imprimo.
	longitudArray := len(frutas)
	fmt.Printf("El tamaño del array es %d → %T\n", longitudArray, len(frutas))

	// Imprimo la capacidad del array.
	fmt.Println("La capacidad del array es: ", cap(frutas))

	// Declaro y defino un array de nombres.
	alumnos := [...]string{"Roberto", "Michel", "Seyi", "Diego"}

	// Muestro los datos del array, utilizando un bucle for
	for i := 0; i < len(alumnos); i++ {
		fmt.Printf(`
Índice: %d
Valor: %s
		`, i, alumnos[i])
	}

	// Recorremos el array, utilizando [range]
	for indice, valor := range alumnos {
		fmt.Printf(`
Índice	→ %d
Valor		→ %s
		`, indice, valor)
		fmt.Println(cap(alumnos))
	}
}
