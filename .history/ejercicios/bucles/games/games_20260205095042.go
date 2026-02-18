package games

import (
	c "bucles/config"
	g "bucles/games"
	v "bucles/view"
)

func GetGame(optionSelected int, parametersGame ...any) func() {
	switch optionSelected {
	case 1, 2, 3, 4:
		v.ShowMockText(c.SELECTED, optionSelected)
	case 5:
		return func() {}()
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
