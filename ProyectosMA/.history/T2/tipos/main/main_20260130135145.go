package main

func main() {
	fmt.Println("Integer Types:")
	PrintIntRange("int8", 8, true)
	PrintIntRange("uint8", 8, false)
	PrintIntRange("int16", 16, true)
	PrintIntRange("uint16", 16, false)
	PrintIntRange("int32", 32, true)
	PrintIntRange("uint32", 32, false)
	PrintIntRange("int64", 64, true)
	PrintIntRange("uint64", 64, false)

	fmt.Println("\nFloating Types:")
	printFloatRange("float32", 32)
	printFloatRange("float64", 64)

	fmt.Println("\nOther Info:")
	fmt.Printf("Type of math.MaxFloat64: %s\n", reflect.TypeOf(math.MaxFloat64))

	fmt.Println("valor: %d. Type %T", any, any)

	fmt.Println("valor: %d. Type %T")
}
