module ProyectoGo/ejercicios

go 1.25.6

package main

import (
	"ProyectoGo/ejercicios"
	"fmt"
)

// Para importar varios paquetes, sepárelos con paréntesis.
// import (
//     "fmt"
//     "os"
// )

// La función main es el punto de entrada de un programa Go.
func main() {
	fmt.Println("Hello, World! 🌍")
	fmt.Println("Miguel Angel\nMedina")
	ejercicios.MostrarNombre()

}
