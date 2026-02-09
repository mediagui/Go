package slices

import "fmt"

func Tratarslices() {
	fmt.Println("### slices ###")
	/*
		Son parecidos a los arrays, que también almacenan datos, pero los slices lamacenan una cantidad de datos indeterminados.
	*/

	// Declaro un [slices]
	var rebanadaUno []string

	// Añado datos.
	rebanadaUno = append(rebanadaUno, "dato 1", "dato2", "dato 3", "dato n")

	fmt.Println(rebanadaUno)

	// Creo un [slice] de números.
	numeros := []int{10, 20, 30}

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
	ingredientes = append(ingredientes, "Lechuga", "Tomate", "Apio", "Puerro")

	receta := copy(colores, ingredientes)

	fmt.Println(receta)
}
