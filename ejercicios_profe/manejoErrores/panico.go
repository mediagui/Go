// Paquete del proyecto.
package manejoErrores

// Importación de paquetes
import "fmt"

// "os" // Paquete que trabaja con archivos y el sistema operativo

/*
La palabra reservada [panic] se utiliza para generar la interrupción inmediata en la ejecución normal del programa. Cuando se produce un [panic], el programa se detiene y comienza a desenrrollar (unwind) la pila de llamadas, ejecutando cualquier función [defer] (diferida) en el camino.

Un [panic] se puede provocar intencionalmente utilizando la función [panic()] o puede ocurrir automáticamente en respuesta a una situación excepcional, como una división entre 0.
*/

func ExplPanic() {
	divideDos(100, 10)
	divideDos(200, 25)
	divideDos(34, 0)
	divideDos(124, 8)
}

// Función que recive un dividendo y divisor y devuelve el resultado de la división.
func divideDos(dividendo, divisor int) {
	// Llamo a la función que comprueba si el número que esta dividiendo es 0.
	/* validarZero(divisor)
	fmt.Println(dividendo / divisor) */
	/*
		Con [recover], podemos captruar y manejar un panic. Cuando lo utilizamos en una función [defer], capturaremos cualquier [panic] que ocurra en esa función o en las funciones llamadas desde ella. Esto permite al programa recuperar el control y realizar acciones adicionales antes de que la ejecución finalice.

		Para explicar esto, voy a utilizar una función "anónima". Para crear una función anónima en go, lo haremos de la siguiente manera:

		func(){
			// Bloque de instrucciones
		}()
	*/
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r) // Imprimo el valor recuperado en caso de un [panic]
		}
	}()

	validarZero(divisor)
	fmt.Println(dividendo / divisor)
}

func validarZero(divisor int) {
	if divisor == 0 {
		panic("No puedes dividir por cero")
	}
}
