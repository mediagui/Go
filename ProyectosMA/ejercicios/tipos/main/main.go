package main

import (
	"fmt"
	"math"
	"reflect"
	"tipos"
)

func main() {
	fmt.Println("Integer Types:")
	tipos.PrintIntRange("int8", 8, true)
	tipos.PrintIntRange("uint8", 8, false)
	tipos.PrintIntRange("int16", 16, true)
	tipos.PrintIntRange("uint16", 16, false)
	tipos.PrintIntRange("int32", 32, true)
	tipos.PrintIntRange("uint32", 32, false)
	tipos.PrintIntRange("int64", 64, true)
	tipos.PrintIntRange("uint64", 64, false)

	fmt.Println("\nFloating Types:")
	tipos.PrintFloatRange("float32", 32)
	tipos.PrintFloatRange("float64", 64)

	fmt.Println("\nOther Info:")
	fmt.Printf("Type of math.MaxFloat64: %s\n", reflect.TypeOf(math.MaxFloat64))

	fmt.Printf("Boolean %v", true)
}
