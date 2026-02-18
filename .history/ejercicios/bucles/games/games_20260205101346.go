package games

import (
	c "bucles/config"
	g "bucles/games"
	v "bucles/view"
)

func GetGame(optionSelected int, parametersGame any) func any) any {
	switch optionSelected {
	case 1, 2, 3, 4:
		v.ShowMockText(c.SELECTED)
	case 5:

	return func(x any) any { n := x.(int) res := 1 for i := 2; i <= n; i++ { res *= i } return res }
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
