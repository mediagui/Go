// Vamos a programar el sistema de puntos de una nueva máquina arcade.
// Necesitamos funciones que sumen puntos, apliquen multiplicadores (Poer-ups) y calculen “combos” especiales.

// Instrucciones del código
// Se debe escribir un único archivo [main.go] que contenga:

//! 1.    El Combo Especial (Recursividad)
// Crea una función llamada [SumaCombo(n int) int]
// * Si un jugador hace un combo de 5 golpes, los puntos son 5 + 4 + 3 + 2 +1.
// * Regla: La función debe sumarse a sí misma hasta llegar a 1.

//! 2.    El procesador de la partida (Orden superior + Variádica)
// Crear una función [RegistrarPuntos]
// * Variadica: Debe recibir muchos números (las puntuaciones de varias “partidas”)
// * Orden Superior: Debe recibir una función que decida si esos puntos se quedan igual o cambian (un “modificador”).
// * Retorno: Imprime el resultado final.

//! 3.    El “Power-Up” (Clousure)
// Crear una función llamada [NuevoMultiplicador(factor int)]
// * Debe devolver una función anónima que multiplique cualquier número por ese factor.
// * Esto sirve para crear objetos como el “Doble Puntuación” o “Triple Puntuación”

// Lo que debe pasar en el [main]
// Para que el ejercicio esté completo, se debe:
// * Usar una función anónima dentro de [RegistrarPuntos] para restar 10 puntos de “penalización” a una serie de puntuaciones.
// * Crear un multiplicador de “Doble Puntos” usando el clousure.
// * Calcular cuánto vale un “Como de 10” usando la función recursiva.

package main

// Lo que debe pasar en el [main]
// Para que el ejercicio esté completo, se debe:
// * Usar una función anónima dentro de [RegistrarPuntos] para restar 10 puntos de “penalización” a una serie de puntuaciones.
// * Crear un multiplicador de “Doble Puntos” usando el clousure.
// * Calcular cuánto vale un “Combo de 10” usando la función recursiva.
func main() {

	RegistrarPuntos(func(puntos int) int {
		return puntos - 10
	}, 100, 200, 300)

	multiplicador := NuevoMultiplicador(2)

	println("Doble puntos:", multiplicador(50))

	println("Combo de 10:", SumaCombo(10))
}

// ! 1.    El Combo Especial (Recursividad)
// Crea una función llamada [SumaCombo(n int) int]
// * Si un jugador hace un combo de 5 golpes, los puntos son 5 + 4 + 3 + 2 +1.
// * Regla: La función debe sumarse a sí misma hasta llegar a 1.
func SumaCombo(n int) int {

	if n == 1 {
		return 1
	}
	return n + SumaCombo(n-1)
}

// ! 2.    El procesador de la partida (Orden superior + Variádica)
// Crear una función [RegistrarPuntos]
// * Variadica: Debe recibir muchos números (las puntuaciones de varias “partidas”)
// * Orden Superior: Debe recibir una función que decida si esos puntos se quedan igual o cambian (un “modificador”).
// * Retorno: Imprime el resultado final.
func RegistrarPuntos(modificador func(int) int, puntos ...int) {
	for _, p := range puntos {
		println("Puntos registrados:", modificador(p))
	}
}

// ! 3.    El “Power-Up” (Clousure)
// Crear una función llamada [NuevoMultiplicador(factor int)]
// * Debe devolver una función anónima que multiplique cualquier número por ese factor.
// * Esto sirve para crear objetos como el “Doble Puntuación” o “Triple Puntuación”
func NuevoMultiplicador(factor int) func(int) int {
	return func(puntos int) int {
		return puntos * factor
	}
}
