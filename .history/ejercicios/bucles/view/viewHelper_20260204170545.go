// This package holds all the logic thar regards to the console view
package view

import (
	"bucles/util"
	"fmt"
)

// Shows the entry menu in the console
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

// Read the selected option entered
func ReadOptionFromConsole() int {

	var selectedOption int
	fmt.Scanln(&selectedOption)

	return selectedOption
}

// Function to wait for the user to press enter before continue with the execution of the program, this is useful to show the result of a game before clear the console and show the menu again.
func PressEnterToContinue() {

	fmt.Print("Presiona Enter para continuar...")
	fmt.Scanln()
}

func ExitGamingSession() {
	fmt.Println("")
}
