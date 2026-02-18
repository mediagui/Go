// Paquete del proyecto.
package manejoErrores

import (
	"fmt"
	"strconv"
)

func ErroresBasicos() {
	fmt.Println("Funciono ErroresBasicos")
	// Declaro y defino una variable de tipo string y almaceno un numero.
	// str := "123" // Esta conversión NO daría error
	str := "123g" // Esta conversión daría error

	// Capturo la converión a número. Esto devuelve el valor y puede devolver un error.
	num, err := strconv.Atoi(str)

	if err != nil {
		// Al no poder realizarse la conversión, se devuelve el valor por defecto del [int], que es 0.
		fmt.Println("Numero: ", num, " Error: ", err)
		return
	} else {
		fmt.Println("Número: ", num)
	}
}
