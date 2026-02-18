package slicen

import "fmt"

func TratarSlicen() {
	fmt.Println("### SLICEN ###")
	/*
		Son parecidos a los arrays, que también almacenan datos, pero los slicen lamacenan una cantidad de datos indeterminados.
	*/

	// Declaro un [slicen]
	var rebanadaUno []string

	// Añado datos.
	rebanadaUno = append(rebanadaUno, "dato 1", "dato2", "dato 3", "dato n")

	fmt.Println(rebanadaUno)

	// Creo un [slice] de números.
	numeros := []int{10, 20, 30, 40}

	// Imprimo el listado de números.
	for indice, valor := range numeros {
		fmt.Printf("El valor para el índice [%d] es [%d]\n", indice, valor)
	}

	// Obtengo  la longitud del slice.
	fmt.Println("Longitud → ", len(numeros))

	// Obtengo la capacidad de slice.
	fmt.Println("Capacidad → ", cap(numeros))

	/*
		Con [make()], puedo crear un [slice], reservando un tamaño y una capacidad máxima.
	*/
	// Reservar longitud desde el inicio.
	edades := make([]int, 3) // Slice de longitud 3: []
	// Imprimo el contenido del slice de edades.
	fmt.Println(edades)
	fmt.Println(len(edades))

	// Reservar longitud y capacidad desde el inicio.
	edadesDos := make([]int, 3, 10)
	fmt.Println(edadesDos)
	edadesDos = append(edadesDos, 8)
	fmt.Println(edadesDos)
	fmt.Println(len(edadesDos))
	fmt.Println(cap(edadesDos))

	// Recojo un sub fragmento de un [slice].
	colores := []string{"rojo", "azul", "verde", "amarillo", "dorado", "celeste", "cyan", "magente"}

	// Bucle para recorrer el slice de colores.
	for indiceColores, color := range colores {
		fmt.Println("El color ["+color+"] tiene asignado el índie [", indiceColores, "]")
	}

	fmt.Println(colores)
	fmt.Println(len(colores))
	fmt.Println(cap(colores))

	fragmento := colores[1:4]
	fmt.Println(fragmento)

	colores = append(colores[:4], colores[5:]...)

	fmt.Println(colores)
	fmt.Println(len(colores))
	fmt.Println(cap(colores))

	/*
		Funciones [make()] y [copy()]
	*/

	// Creo un [slice] con [make()]
	ingredientes := make([]string, 5, 10)

	// Imprimo la información del slice [ingedientes[]]
	fmt.Println(len(ingredientes))
	fmt.Println(cap(ingredientes))
	// Añado elementos a partir del último ocupado.
	ingredientes = append(ingredientes, "Lechuga", "Tomate", "Apio", "Puerro")
	fmt.Println(colores)
	// Como vemos al imprimirlo, se imprimen 5 espacios en blanco al principio.
	fmt.Println(ingredientes)

	/*
		[copy()]

		#sintaxis#
		copy(destino, origen)

		#acción#
		Copia elementos de un [slicen] a otro. Es importante saber que solo copia tantos elementos como quepan en el slicen más pequeño.

		#devolución#
		Devuelve un entero que representa la cantidad de elemntos que se han copiado
	*/
	receta := copy(colores, ingredientes) /*
		En este caso, se copian los elementos de [ingredientes] hacia [colores]. Se copiarán tantos como quepan en [colores], puesto que ahora tiene 7 elementos e [ingredientes] tiene 9. Al tener los 5 primeros como elmentos vacíos (valor por defecto de los string), es lo que almacena.
	*/
	for indice, valorColes2 := range colores {
		fmt.Printf("El valor para el ínice [%d] es [%s] \n", indice, valorColes2)
	}

	// Imprime la cantidad de elementos copiados.
	fmt.Println(receta)
}
