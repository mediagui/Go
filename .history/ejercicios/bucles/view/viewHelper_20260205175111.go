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

	1. Guess the letter
	2. Print Odd or Even
	3. Factorial calculation
	4. Sum without reach 50
	5. Areas calculation
	6. End

	   Choose: `

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

func ShowMockText(texToShow ...string) {
	fmt.Println(texToShow)
}

func ResquestValue(question string) string {
	var answer string
	fmt.Print(question)
	fmt.Scanln(&answer)

	return answer

}

// This package holds all the logic thar regards to the console view
package view

import (
	c "bucles/config"
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

	fmt.Print(c.PRESS_ENTER)
	fmt.Scanln()
}

func ShowExitGamingSessionMessage() {
	fmt.Println(c.END_MESSAGE)
}

func ShowInvalidOptionMessage() {
	fmt.Println(c.INVALID_OPTION_MESSAGE)

}
