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
	"bucles/config"
	"bucles/games"
	"bucles/util/validation"
	menu "bucles/view"

	"fmt"
)

// ============================================================================
// FUNCIÓN MAIN
// ============================================================================
//
// main es el punto de entrada principal del programa.
//
// FLUJO DE EJECUCIÓN:
// 1. Inicia un bucle infinito (for {})
// 2. Muestra el menú de opciones al usuario
// 3. Lee la opción seleccionada por el usuario
// 4. Valida que la opción sea válida
// 5. Si es válida, ejecuta playGame() con esa opción
// 6. Espera a que el usuario presione Enter para continuar
// 7. Repite desde el paso 2
//
// ESTRUCTURA DEL BUCLE:
//
//	for {
//	    - Mostrar menú
//	    - Leer opción
//	    - Validar opción
//	    - Ejecutar juego si es válida
//	    - Esperar Enter
//	}
//
// Este patrón crea una experiencia interactiva donde el usuario puede
// elegir diferentes ejercicios repetidamente.
func main() {

	for {
		// Muestra las opciones disponibles al usuario
		menu.ShowMenuInConsole()

		// Lee la opción seleccionada desde la entrada estándar
		selectedOption := menu.ReadOptionFromConsole()

		// Valida que la opción sea un número válido en el rango esperado
		if validation.IsAValidOption(selectedOption) {
			// Si es válida, ejecuta el juego/ejercicio correspondiente
			playGame(selectedOption)

		} else {
			// Si no es válida, simplemente continúa el bucle sin hacer nada
			// (el usuario verá un mensaje de error en el menú siguiente)
		}

		// Espera a que el usuario presione Enter antes de mostrar el menú nuevamente
		menu.PressEnterToContinue()

		menu.PressEnterToContinue()
	}

}

// ============================================================================
// FUNCIÓN PLAYGAME
// ============================================================================
//
// playGame es la función que ejecuta el juego/ejercicio seleccionado.
// Demuestra el uso práctico de la interfaz GameFunction y el patrón
// de funciones heterogéneas implementado en games.go.
//
// PARÁMETRO:
// - selectedOption (int): El número de la opción seleccionada por el usuario
//
// FLUJO:
// 1. Obtiene la función encapsulada llamando a games.GetGame()
// 2. Ejecuta la función llamando a game.Execute(nil)
// 3. Muestra el resultado en la consola
// 4. Muestra la opción que se ejecutó
//
// EJEMPLO DE EJECUCIÓN:
//
//	selectedOption = 1
//	game = games.GetGame(1)      // Retorna interfaz GameFunction con función de factorial
//	result = game.Execute(nil)   // Ejecuta la función y retorna el resultado
//	Imprime: "Resultado: 120" (si el usuario ingresó 5)
//	        "Opción seleccionada: 1"
//
// VENTAJAS DE ESTE DISEÑO:
// - playGame() no necesita saber la firma específica de cada función
// - No hay condicionales para manejar diferentes tipos de retorno
// - Es extensible: agregar nuevos ejercicios solo requiere modificar GetGame()
// - La interfaz GameFunction proporciona un contrato universal
//
// EJEMPLO DE CÓMO AGREGAR UN NUEVO EJERCICIO:
//  1. En games.go, agregar un nuevo case en GetGame():
//     case 6:
//     return buildFunction(func(p any) any {
//     // Código del nuevo ejercicio
//     return resultado
//     })
//
// 2. No es necesario modificar playGame() ni main()
// 3. El nuevo ejercicio funcionará automáticamente con el flujo existente
func playGame(selectedOption int) {

	// Obtiene la función encapsulada basada en la opción seleccionada
	// El retorno es siempre de tipo GameFunction (interfaz)
	// Aunque internamente contiene funciones con diferentes firmas
	game := games.GetGame(selectedOption)

	// Ejecuta la función encapsulada llamando a su método Execute()
	// Se pasa nil como argumento (no se usa en este caso)
	// El resultado puede ser de cualquier tipo (se retorna como any)
	result := game.Execute(nil)

	// Muestra el resultado obtenido en la consola
	fmt.Println(config.RESULT, result)

	// Muestra la opción que fue ejecutada
	fmt.Println(config.SELECTED, selectedOption)
}
