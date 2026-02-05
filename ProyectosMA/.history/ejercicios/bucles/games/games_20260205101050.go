package games

import (
	c "bucles/config"
	g "bucles/games"
	v "bucles/view"
)

func GetGame(optionSelected int, parametersGame ...any) func(...any) any {
	switch optionSelected {
	case 1, 2, 3, 4:
		v.ShowMockText(c.SELECTED)
	case 5:

		return BuildFunction(func(num int) int {
			if num == 0 || num == 1 {
				return 1
			}
			return int(num) * factorial(int(num)-1)
		})

	default:

	}
}

func BuildFunction[T any, R any](f func(T) R) func(T) R {
	return f
}

func factorial(num int) int {
	result := 1

	if num > 1 {

		return factorial(num - 1)

	}
	return result
}
