package main

import (
	"fmt"
	"github.com/fatih/color"
)

// Para importar varios paquetes, sepárelos con paréntesis.
// import (
//     "fmt"
//     "os"
// )

// La función main es el punto de entrada de un programa Go.
func MostrarNombre() {
	fmt.Println("Miguel Angel\nMedina")

	color.Blue("Color azul")
}

func main() {
	fmt.Println("Hello, World! 🌍")
	MostrarNombre()

}
