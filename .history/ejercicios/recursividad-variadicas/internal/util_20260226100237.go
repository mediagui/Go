package internal

import (
	"math"
	"runtime"
)

func CanCalculateFactorial(n int) {

	// Obtener arquitectura (32 o 64 bits)

	math.MaxInt
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
