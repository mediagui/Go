// Vamos a programar el sistema de puntos de una nueva máquina arcade. Necesitamos funciones que sumen puntos, apliquen multiplicadores (Poer-ups) y calculen “combos” especiales.
// Instrucciones del código
// Se debe escribir un único archivo [main.go] que contenga:

// 1.    El Combo Especial (Recursividad)
// Crea una función llamada [SumaCombo(n int) int]
// v Si un jugador hace un combo de 5 golpes, los puntos son 5 + 4 + 3 + 2 +1.
// v Regla: La función debe sumarse a sí misma hasta llegar a 1.

// 2.    El procesador de la partida (Orden superior + Variádica)
// Crear una función [RegistrarPuntos]
// v Variadica: Debe recibir muchos números (las puntuaciones de varias “partidas”)
// v Orden Superior: Debe recibir una función que decida si esos puntos se quedan igual o cambian (un “modificador”).
// v Retorno: Imprime el resultado final.

// 3.    El “Power-Up” (Clousure)
// Crear una función llamada [NuevoMultiplicador(factor int)]
// v Debe devolver una función anónima que multiplique cualquier número por ese factor.
// v Esto sirve para crear objetos como el “Doble Puntuación” o “Triple Puntuación”

// Lo que debe pasar en el [main]

// Para que el ejercicio esté completo, se debe:

// v Usar una función anónima dentro de [RegistrarPuntos] para restar 10 puntos de “penalización” a una serie de puntuaciones.

// v Crear un multiplicador de “Doble Puntos” usando el clousure.

// v Calcular cuánto vale un “Como de 10” usando la función recursiva.

package main
