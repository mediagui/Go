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
	// menu := "\n1. Adivina la letra\n2. Imprime par/impar\n3. Calcula el factorial\n4. Suma sin pasarte de 50\n5. Cálculo de areas\n6. Terminar\tOpción: "

	fmt.Print(menu)

}
