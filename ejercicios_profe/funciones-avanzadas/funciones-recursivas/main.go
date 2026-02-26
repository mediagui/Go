// Paquete principal.
package main

// Importaciones
import "fmt"

/*
######################
#FUNCIONES RECURSIVAS#
######################

La recursividad es un concepto de la programación en el que una función se llama a sí misma de manera repetitiva hasta que se cumple una condición de salida. Una función recursiva es una función que contiene una llamada a sí misma dentro de su definición.

En Go, las funciones recursivas se definen de manera similar a cualquier otra función, pero se incluye una llamada a sí misma dentro del cuerpo de esta.

func nombreFunción() {
	nombreFunción()
}

*/

// Ejemplo de función recursiva que muestra el factorial de un número.
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

/*
En este ejemplo, la función factorial toma un parámetro entero [num] y devuelve el factorial de [num]. El factorial de un número entero, se define como [num * (num-1) * num * (num-2) * num * (num-n)]

La funciópn factorial utiliza la condición [if] para comprobar si el valor de [num] es 0. Si es así, la función devuelve 1, que es el factorial de 0. Si [num] es diferente de 0, la función devuelve el producto de [num] y el factorial de [num-1]. La lamada a [factorial(num-1)] es la llamada recursiva, que se repite hasta que la condicón [num==0] se cumple y devuelva 1.
*/

func main() {
	// Imprimo el factorial del número.
	fmt.Println(factorial(5))
}
