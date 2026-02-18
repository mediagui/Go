// This package holds the the logic to validate selected options.
// The values available to choose are defined within the LOWER_LIMIT and UPPER_LIMIT constants in the config package.
package validation

import (
	c "bucles/config"
	"fmt"
)

// Returns true if the selected option value is in between the values defined in LOWER_LIMIT and UPPER_LIMIT
func IsAValidOption(selectedOption int) bool {
	lowerLimitOk := c.LOWER_LIMIT >= selectedOption
	upperLimitOk := c.UPPER_LIMIT >= selectedOption

	return upperLimitOk && lowerLimitOk

}

func IsExitSelected(selectedOption int) bool {

	return selectedOption == c.UPPER_LIMIT

}

func CheckError(err error, message string) {
	if err != nil {
		fmt.Println(message)
	}

}
