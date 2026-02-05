package games

func GetGame(optionSelected int) func {
	switch optionSelected {
	case 1,2,3,4:
		view.ShowMockText(c.SELECTED, optionSelected)
	case 5:
		games.
	default:

	}
}

func factorial(num int) int {
	result := 1


	if num >1 {

		return factorial(num -1)

	}
	return result
}


