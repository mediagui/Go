package main

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
