package tipos


import (
	"fmt"
	"math"
	"reflect"
)

// Helper function to print min and max for integer types
func printIntRange(name string, bits int, signed bool) {
	if signed {
		min := -(1 << (bits - 1))
		max := (1 << (bits - 1)) - 1
		fmt.Printf("%-6s | Min: %d | Max: %d\n", name, min, max)
	} else {
		min := 0
		max := (1 << bits) - 1
		fmt.Printf("%-6s | Min: %d | Max: %d\n", name, min, max)
	}
}








	fmt.Println("valor: %d. Type %T", any, any)

	fmt.Println("valor: %d. Type %T")

}

