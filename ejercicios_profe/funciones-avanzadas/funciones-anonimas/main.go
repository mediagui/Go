package main

import "fmt"

/*
####################
#FUNCIONES ANÓNIMAS#
####################

Son las funciones que no tienen un identificador y que pueden ser definidas en el lugar donde son necesitadas, en lugar de ser definidas globalmente. En Go, estas se crean utilizando la palabra clave [func] seguida de los parámetros y el cuerpo de la función entre llaves [{}].

#SINTAXIS#

func (){
	// Bloque de instrucciones
}()
*/

// Función que recibe como parámetro un nombre y una función anónima.
func saludos(nombre string, f func(string)) {
	f(nombre)
}

// Declaro y defino varias funciones que tratarán con números, y que luego serán lamadas como argumento a modo de función anónima.
func duplicar(num int) int {
	return num * 2
}

func triplicar(num int) int {
	return num * 3
}

func aplicar(f func(int) int, num int) int {
	return f(num)
}

func main() {
	// Creo una función anónima que imprime un mensaje de bienvenida.
	func() {
		fmt.Println("Hola, soy una función anónima")
	}()

	// También puedo almacenar el resultado en una variable. Para ello, debo eliminar los [()] del final. Tmabién puedo indicar parámetros y enviárselos en la llamada.
	despedida := func(nombre string) {
		fmt.Printf("Adioooos %s con el corazoooon\n", nombre)
	}

	despedida("Víctorino")
	saludos("Michle", despedida)

	fmt.Println(`
	#Función anónima como argumento#
	`)

	resultado1 := aplicar(duplicar, 5)
	resultado2 := aplicar(triplicar, 5)

	fmt.Println(resultado1, resultado2)
}

/*
	En este ejemplo, se definen dos funciones [duplicar] y [triplicar], que tomam un argumento [num] de tipo [iny] y devuelve n el valor de [num] multiplicado por 2 y 3 respectivamente.

	Luego, se definen una función [aplicar] que toma dos argumentos: [f] de tipo [func(int) int] y [num] de tipo [int]. Dentro de la función, se llama a la función [f] pasando el argumento [n].

	En la función [main()], se llama a la función [aplicar()] dos veces con diferentes funciones y argumentos. El resultado de cada llamada se almacena en una variable y se imprime.
*/
