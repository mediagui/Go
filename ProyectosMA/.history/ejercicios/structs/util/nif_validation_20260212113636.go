package util

import (
	"fmt"
	"regexp"
	"strings"
)

// validateNIF checks if a given string is a valid Spanish NIF
func validateNIF(nif string) bool {
	// Normalize input: remove spaces and convert to uppercase
	nif = strings.ToUpper(strings.TrimSpace(nif))

	// Regular expression: 8 digits + 1 letter
	re := regexp.MustCompile(`^(\d{8})([A-Z])$`)
	matches := re.FindStringSubmatch(nif)
	if matches == nil {
		return false // Format does not match
	}

	// Extract number and letter
	numberPart := matches[1]
	letterPart := matches[2]

	// Sequence of valid letters for NIF
	letters := "TRWAGMYFPDXBNJZSQVHLCKE"

	// Convert number to integer
	var num int
	_, err := fmt.Sscanf(numberPart, "%d", &num)
	if err != nil {
		return false
	}

	// Calculate expected letter
	expectedLetter := string(letters[num%23])

	// Compare with provided letter
	return letterPart == expectedLetter
}
