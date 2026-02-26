package main

import "fmt"

/*
#############################
#FUNCIONES DE ORDEN SUPERIOR#
#############################

Son funciones que aceptan otra función como argumento y/o devuelve una función como resultado. Esto permite una programación más modular y flexible, ya que las funciones de orden superior pueden utilizarse para crear funciones personalizadas en tiempo de ejecución.

Se utilizan para implementar patrones de programación comunes, como el mapeo, filtrado y reducción de datos en colecciones. Por ejemplo, la función [map()] es una función de orden superior que toma una función como argumento y la aplica a cada elemento de una colección, devolviendo una nueva colección con los resultados de la aplicación de l afunción a cada elemento.
*/

// Función que recibe a otr afunción como argumento.
func doble(f func(int) int, x int) int {
	return f(x * 2)
}

func sumarUno(x int) int {
	return x + 1
}

func main() {
	// Almaceno en una variable el resultado de la función "superior", pasándole como argumento la función que suma 1.

	r := doble(sumarUno, 8)
	fmt.Println("Resultado: ", r)
}
