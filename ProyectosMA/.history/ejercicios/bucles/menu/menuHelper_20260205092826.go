// This package holds all the logic thar regards to the console view
package menu

import (
	c "bucles/config"
	"bucles/util"

	"fmt"
)

func ShowMenuInConsole() {

	util.ClearConsole()

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

func ReadOptionFromConsole() int {

	var selectedOption int
	fmt.Scanln(&selectedOption)

	return selectedOption
}

func PressEnterToContinue() {
	fmt.Print(c.PRESS_ENTER_TO_CONTINUE)
	fmt.Scanln()
}
