package tipos

import "fmt"

func MostrarValorDefault() {
	var (
		defaultInt    int
		defaultUint   uint
		defaultFloat  float32
		defaultBool   bool
		defaultString string
	)

	fmt.Println(defaultInt)    // 0
	fmt.Println(defaultUint)   // 0
	fmt.Println(defaultFloat)  // 0
	fmt.Println(defaultBool)   // false
	fmt.Println(defaultString) // Caracter vacío " "
}
