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

// Helper function to print min and max for floating types
func printFloatRange(name string, bits int) {
	if bits == 32 {
		fmt.Printf("%-6s | Min: %g | Max: %g\n", name, -math.MaxFloat32, math.MaxFloat32)
	} else if bits == 64 {
		fmt.Printf("%-6s | Min: %g | Max: %g\n", name, -math.MaxFloat64, math.MaxFloat64)
	}
}

func main() {
	fmt.Println("Integer Types:")
	printIntRange("int8", 8, true)
	printIntRange("uint8", 8, false)
	printIntRange("int16", 16, true)
	printIntRange("uint16", 16, false)
	printIntRange("int32", 32, true)
	printIntRange("uint32", 32, false)
	printIntRange("int64", 64, true)
	printIntRange("uint64", 64, false)

	fmt.Println("\nFloating Types:")
	printFloatRange("float32", 32)
	printFloatRange("float64", 64)

	fmt.Println("\nOther Info:")
	fmt.Printf("Type of math.MaxFloat64: %s\n", reflect.TypeOf(math.MaxFloat64))
}




	fmt.Println("valor: %d. Type %T", any, any)

	fmt.Println("valor: %d. Type %T")

}

