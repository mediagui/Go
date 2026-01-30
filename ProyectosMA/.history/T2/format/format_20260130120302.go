package format

import (
	"fmt"
	"strconv"
)

func PrintHeader() {
	fmt.Println("ID\tName\t\tPrice\n--\t----------\t---------")
}

func PrintFormattedPrize(prize float64) string {
	return strconv.FormatFloat(prize, 'f', 2, 32)
}
