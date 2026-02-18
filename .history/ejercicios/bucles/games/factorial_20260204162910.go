package games

//
func FactorialOf(number int) int {

	if number ==0 || number == 1 {
		return 1
	} else {
		return number * FactorialOf(number-1)
	}
}
