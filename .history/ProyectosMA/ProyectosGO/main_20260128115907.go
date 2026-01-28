package main

import (
	"fmt"

	"rsc.io/quote"
)

// Para importar varios paquetes, sepárelos con paréntesis.
// import (
//     "fmt"
//     "os"
// )

// La función main es el punto de entrada de un programa Go.
func MostrarNombre() {
	fmt.Println("Miguel Angel\nMedina")
	fmt.Println(quote.Go())
}

func main() {
	fmt.Println("Hello, World! 🌍")
	MostrarNombre()

}
