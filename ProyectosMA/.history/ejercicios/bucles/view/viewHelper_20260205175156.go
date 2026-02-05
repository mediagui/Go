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

func ShowExitGamingSessionMessage() {
	fmt.Println(c.END_MESSAGE)
}

func ShowInvalidOptionMessage() {
	fmt.Println(c.INVALID_OPTION_MESSAGE)

}
