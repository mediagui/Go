//	1.Crea un sistema que pueda adivinar un carácter insertado es una vocal o no. Para eso tiene que ser insertado un carácter por consola.
//	2.Solicita un número e imprime todos los números pares e impares desde 1 hasta ese número con el  mensaje "es par" o
// "es impar" si el número es 5 el resultado será: 1 - es impar 2 - es par 3 – es impar 4 - es par 5 - es impar.
//	3.Escriba un programa que pida un número entero mayor que cero y calcule su factorial.
// 		La factorial es el resultado de multiplicar ese número por sus anteriores hasta la unidad.
//	4.Escribe un programa que permita ir introduciendo una serie indeterminada de números mientras su suma no supere 50.
// 		Cuando esto ocurra, se debe mostrar el total acumulado y el contador de cuantos números se han introducido
//	5.Escribe un programa con un bucle infinito con opciones para elegir, que pueda calcular el área de 2 figuras geométricas, triángulo y rectángulo. En primer lugar, pregunta de qué figura se quiere calcular el área, después solicita los datos que necesites para calcularlo.
//		1) triángulo = b * h/2
//		2) rectángulo = b * h
//		3) Salir: Solo cuando escojas esta opción el bucle se detendrá

package main

import (
	c "bucles/config"
	v "bucles/util/validation"
	menu "bucles/view"

	"fmt"
)

func main() {

	for {

		menu.ShowMenuInConsole()

		selectedOption := menu.ReadOptionFromConsole()

		if isValid := !v.IsAValidOption(selectedOption); isValid {

			playGame(selectedOption)

		} else {

		}

		menu.PressEnterToContinue()

	}

}

func playGame(selectedOption int) {
	fmt.Println(c.SELECTED, selectedOption)
}
