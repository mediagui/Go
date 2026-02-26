package main

import (
	"math"
	"runtime"
)

func CanCalculateFactorial(n int) bool {

	// Obtener arquitectura (32 o 64 bits)

	bits := getIntBitsByArchitecture()

	return n < int(math.Pow(2, float64(bits)) - 1)
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
