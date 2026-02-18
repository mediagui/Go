package games

import (
	c "bucles/config"
	v "bucles/view"
)

func GetGame(optionSelected int, parametersGame any) func(any) any {

	var returnedFunc func(any) any

	switch optionSelected {
	case 1:

		returnedFunc = func(x any) any {
			n := x.(int)
			r := 1
			for i := 2; i <= n; i++ {
				r *= i
			}
			return r
		}

	case 2, 3, 4, 5:
		v.ShowMockText(c.SELECTED)
		fallthrough
	default:
		returnedFunc = func(x any) any {
			return nil
		}
	}
	return returnedFunc
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
