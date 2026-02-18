package main

func main() {
	panic("unimplemented")
}

// Helper function to print min and max for floating types
func printFloatRange(name string, bits int) {
	if bits == 32 {
		fmt.Printf("%-6s | Min: %g | Max: %g\n", name, -math.MaxFloat32, math.MaxFloat32)
	} else if bits == 64 {
		fmt.Printf("%-6s | Min: %g | Max: %g\n", name, -math.MaxFloat64, math.MaxFloat64)
	}
}
