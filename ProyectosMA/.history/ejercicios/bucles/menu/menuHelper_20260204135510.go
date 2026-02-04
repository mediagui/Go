package menu

import "fmt"

func ShowMenu() {

	menu := `
	1. Adivina la letra
	2. Imprime par/impar
	3. Calcula el factorial
	4. Suma sin pasarte de 50
	5. Cálculo de areas
	6. Terminar
	   Opción: `

	fmt.Print(menu)
}

func GetSelectedOption() int {
	var selectedOption int
	fmt.Scanln(&selectedOption)

	return selectedOption
}
