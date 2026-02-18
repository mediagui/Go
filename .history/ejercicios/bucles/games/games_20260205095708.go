package games

import (
	c "bucles/config"
	g "bucles/games"
	v "bucles/view"
)

func GetGame(optionSelected int, parametersGame ...any) func(int) int {
	switch optionSelected {
	case 1, 2, 3, 4:
		v.ShowMockText(c.SELECTED)
	case 5:

		factorial := func(num int) int {
			if num == 0 || num == 1 {
				return 1
			}
			return num * factorial(num-1)
		}

		return factorial
		// return func(parameters ...any) {

		// }(parametersGame)
	default:

	}
}

func factorial(num int) int {
	result := 1

	if num > 1 {

		return factorial(num - 1)

	}
	return result
}
