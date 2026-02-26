package internal

import "runtime"

func CanCalculateFactorial(n int) {

	// Obtener arquitectura
	if runtime.GOARCH == "386" || runtime.GOARCH == "arm" {
		if n > 12 {
			panic("Número demasiado grande para calcular factorial en esta arquitectura")
		}
	}
}
