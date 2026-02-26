package internal

import (
	"math"
	"runtime"
)

func canCalculateFactorial(n int) bool {

	// Obtener arquitectura (32 o 64 bits)

	bits := getIntBitsByArchitecture()

	return n < int(math.Pow(2, float64(bits))-1)
}

// Calculo del factorial usando recursividad
func factorial(numero int) (int, error) {

	if numero == 0 {
		return 1, nil
	}

	if !canCalculateFactorial(numero) {
		panic("Número demasiado grande para calcular factorial en esta arquitectura")
	}

	result, err := factorial(numero - 1)
	if err != nil {
		return 0, err
	}

	return numero * result, nil
}

func UsoDeFuncAnonima() {

	println("Usamos una funcion anónima para calcular el cuadrado de un número")

	// Función anónima que calcula el cuadrado de un número
	cuadrado := func(numero int) int {
		return numero * numero
	}

	// Llamada a la función anónima
	resultado1 := cuadrado(5)
	println("Cuadrado de 5:", resultado1)

}

func UsoDeFuncionVariadicaYAnonima() {

	func(numeros ...int) {
		suma := 0
		for _, n := range numeros {
			suma += n
		}
		println("Suma de los números [1, 2, 3, 4, 5]:", suma)
	}(1, 2, 3, 4, 5)

}

func CalculoDeFactorial() {
	println("Calculamos el factorial de 10.000.")
	valor, err := factorial(10_000)
	if err != nil {
		println("Error calculando factorial:", err)
		return
	}

	valor, err = factorial(100)
	if err != nil {
		println("Error calculando factorial:", err)
		return
	}

	println("Resultado:", valor)
}

func getIntBitsByArchitecture() int {
	architecture := runtime.GOARCH
	switch architecture {
	case "386", "arm":
		return 32
	default:
		return 64
	}
}
