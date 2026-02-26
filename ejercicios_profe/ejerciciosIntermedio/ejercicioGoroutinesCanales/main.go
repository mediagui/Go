// Indico el paqute al que pertenezco.
package main

/*
###############
#IMPORTACIONES#
###############
*/
import "fmt"

// Función que suma un rango y devuelve el resultado final.
func sumarRango(inicio, fin int, ch chan int) {
	suma := 0
	for i := inicio; i <= fin; i++ {
		fmt.Printf("El resultado de sumar %d + %d es: %d\n", suma, i, suma+i)
		suma += i
	}
	ch <- suma // Envío al canal el resultado.
}

func main() {
	// Instancio el canal.
	ch := make(chan int)

	// Lanzo las tres goroutines.
	go sumarRango(1, 5, ch)
	go sumarRango(6, 10, ch)
	go sumarRango(11, 15, ch)

	// Recibo los tres resultados.
	total := 0
	for i := 0; i < 3; i++ {
		total += <-ch
	}
	// Imprimo el resultado de la suma total.
	fmt.Println("La suma total es: ", total)
}
