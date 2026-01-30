package format

import (
	"fmt"
	"strconv"
)

func PrintHeader() {
	fmt.Println("ID\tName\t\tPrice\n--\t----------\t---------")
}

func PrintFormattedPrize(prize float32) string {
	return strconv.FormatFloat(float64(prize), 'f', 2, 32)
}

func PrintInvoiceLine(line)
