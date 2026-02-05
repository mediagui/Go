package games

func factorial(num int) int {
	result := 1


	if num >1 {

		return factorial(num -1)

	}
	return result
}
