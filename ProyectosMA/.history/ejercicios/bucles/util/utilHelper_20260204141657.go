// This package holds all the method and function utilities.
package util

import "fmt"

// Cleans the console as the clear or cls command do.
func ClearConsole() {
	fmt.Print("\033[H\033[2J")
}
