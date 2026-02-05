package games

import (
	c "bucles/config"
	v "bucles/util/validation"
	menu "bucles/view"

	"fmt"
	"slices"
	"strconv"
)

// ============================================================================
// INTERFAZ GAMFUNCTION
// ============================================================================
//
// GameFunction es una interfaz que actúa como contrato común para encapsular
// funciones con diferentes firmas (tipos de entrada y salida variables).
//
// PROBLEMA RESUELTO:
//   - En Go no se pueden devolver funciones con tipos heterogéneos directamente
//     (ej: func(int)int en un caso y func(string)bool en otro)
//   - La solución es usar una interfaz que defina un método común Execute()
//     que pueda ser implementado por diferentes tipos de wrappers genéricos
//
// VENTAJAS:
// - Permite retornar funciones con diferentes firmas
// - Mantiene type safety mediante genéricos
// - Proporciona una interfaz uniforme para llamar cualquier función
//
// EJEMPLO DE USO:
//
//	game := GetGame(1)           // Retorna una GameFunction
//	resultado := game.Execute()  // Ejecuta la función encapsulada
type GameFunction interface {
	// Execute ejecuta la función encapsulada con los argumentos proporcionados
	// Parámetros:
	//   args ...any: Argumentos variádicos de tipo any (pueden omitirse)
	// Retorna:
	//   any: El resultado de la ejecución (puede ser de cualquier tipo)
	Execute(args ...any) any
}

// ============================================================================
// ESTRUCTURA FUNCTIONWRAPPER
// ============================================================================
//
// FunctionWrapper[T, R any] es un wrapper genérico que encapsula una función
// con tipos específicos (T = tipo entrada, R = tipo retorno).
//
// CÓMO FUNCIONA:
// - T: Tipo genérico del parámetro de entrada
// - R: Tipo genérico del valor de retorno
// - fn: La función original que será envuelta
//
// EJEMPLO:
//
//	FunctionWrapper[int, int] encapsula func(int) int
//	FunctionWrapper[string, bool] encapsula func(string) bool
//	FunctionWrapper[any, any] encapsula func(any) any
type FunctionWrapper[T any, R any] struct {
	// fn es la función original que será ejecutada cuando se llame Execute()
	fn func(T) R
}

// ============================================================================
// MÉTODO EXECUTE - IMPLEMENTA LA INTERFAZ GAMEFUNCTION
// ============================================================================
//
// Execute implementa el método required de la interfaz GameFunction.
// Convierte argumentos dinámicos (any) al tipo específico esperado (T)
// y retorna el resultado de ejecutar la función envuelta.
//
// FLUJO:
// 1. Crea una variable de tipo T con su valor por defecto
// 2. Si hay argumentos y el primero no es nil, lo convierte a tipo T
// 3. Ejecuta la función fn con el argumento convertido
// 4. Retorna el resultado (que será de tipo R, pero como any)
//
// MANEJO DE TIPOS:
// - Si args[0] es nil: usa el valor por defecto del tipo T
// - Si args[0] tiene valor: intenta hacer type assertion a T
//
// EJEMPLO:
//
//	wrapper := &FunctionWrapper[int, int]{fn: func(x int) int { return x * 2 }}
//	resultado := wrapper.Execute(5)  // Retorna 10
//	resultado := wrapper.Execute()   // Retorna 0 (valor por defecto de int)
func (fw *FunctionWrapper[T, R]) Execute(args ...any) any {
	// Inicializa la variable con el valor zero value del tipo T
	var arg T

	// Si se proporcionó al menos un argumento y no es nil,
	// lo convierte al tipo T esperado
	if len(args) > 0 && args[0] != nil {
		arg = args[0].(T) // Type assertion: convierte any a T
	}

	// Ejecuta la función original con el argumento y retorna el resultado
	return fw.fn(arg)
}

// ============================================================================
// FUNCIÓN BUILDFUNCTION
// ============================================================================
//
// buildFunction es una función factory que crea wrappers genéricos.
// Toma una función con tipos específicos y la envuelve en una estructura
// que implementa la interfaz GameFunction.
//
// PARÁMETROS GENÉRICOS:
// - T: Tipo del parámetro de entrada de la función
// - R: Tipo del valor de retorno de la función
//
// PARÁMETRO:
// - f: La función a envolver
//
// RETORNA:
// - GameFunction: Una interfaz que encapsula la función
//
// PROPÓSITO:
//   - Permite que funciones con diferentes firmas implementen
//     la interfaz GameFunction de forma transparente
//   - Usa genéricos para mantener type safety
//
// EJEMPLO DE USO:
//
//	// Envolver una función que calcula factorial (int -> int)
//	game1 := buildFunction(func(n int) int { ... })
//
//	// Envolver una función que valida texto (string -> bool)
//	game2 := buildFunction(func(s string) bool { ... })
//
//	// Ambas devuelven GameFunction, pero con diferentes tipos internos
func buildFunction[T any, R any](f func(T) R) GameFunction {
	// Crea una nueva instancia de FunctionWrapper que implementa GameFunction
	// La función f se almacena en el campo fn del wrapper
	return &FunctionWrapper[T, R]{fn: f}
}

func GetGame(optionSelected int) GameFunction {

	switch optionSelected {
	case 1: // Guess the letter type
		return buildFunction(func(parametersGame any) any {

			letterInput := menu.ResquestValue(c.LET_ME_KNOW_THE_LETTER)

			// Esto es complicarlo, pero es para demostrar el uso de funciones anónimas y closures
			return func(letter string) string {
				vocales := []string{"a", "e", "i", "o", "u"}

				if _, err := strconv.Atoi(letter); err != nil {
					if slices.Contains(vocales, letterInput) {
						return c.VOWEL
					}
					return c.CONSONANT
				}

				return c.ITS_NOT_A_LETTER

			}(letterInput)

		})
	case 2: // Odd or Even
		return buildFunction(func(parametersGame any) any {

			input, err := strconv.Atoi(menu.ResquestValue(c.LET_ME_KNOW_THE_NUMBER))
			v.CheckError(err, c.INVALID_NUMBER)

			if input%2 == 0 {
				return c.EVEN
			}
			return c.ODD
		})
	case 3: // Factorial calculation
		return buildFunction(func(parametersGame any) any {
			n, _ := strconv.Atoi(menu.ResquestValue(c.LET_ME_KNOW_THE_NUMBER))

			var r int = 1

			for i := 2; i <= n; i++ {
				r *= i
			}

			return r
		})
	case 4: // Sum without reach 50
		return buildFunction(func(parametersGame any) any {
			sum := 0

			for sum < 50 {
				input, err := strconv.Atoi(menu.ResquestValue(c.LET_ME_KNOW_THE_NUMBER))
				v.CheckError(err, c.INVALID_NUMBER)
				sum += input
			}

			return sum
		})
	case 5:
		return buildFunction(func(parametersGame any) any {

			// Print the menu area options and request the user to select one of them
			areaOption := func() int {
				var opt int
				var err error
				for {
					fmt.Println("1. Area of a square\n2. Area of a triangle\n3. End")
					opt, err = strconv.Atoi(menu.ResquestValue(c.LET_ME_KNOW_THE_AREA_TYPE))
					v.CheckError(err, c.INVALID_NUMBER)
					if opt == 3 {
						break
					}
				}
				return opt
			}()

			var base, height, area float64
			var err error

			base, err = strconv.ParseFloat(menu.ResquestValue(c.LET_ME_KNOW_THE_BASE), 64)
			v.CheckError(err, c.INVALID_NUMBER)
			height, err = strconv.ParseFloat(menu.ResquestValue(c.LET_ME_KNOW_THE_HEIGHT), 64)
			v.CheckError(err, c.INVALID_NUMBER)

			switch areaOption {
			case 1:
				area = base * height
			case 2:
				area = (base * height) / 2
			default:
				area = 0.0
			}

			return area
		})

	default:
		return buildFunction(func(parametersGame any) any {
			return nil
		})
	}

}
