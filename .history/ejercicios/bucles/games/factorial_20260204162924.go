package games

// Fuction to calculate the factorial of a number, it receives an integer and returns an integer, the result of the factorial of the number
func FactorialOf(number int) int {

	if number ==0 || number == 1 {
		return 1
	} else {
		return number * FactorialOf(number-1)
	}
}
