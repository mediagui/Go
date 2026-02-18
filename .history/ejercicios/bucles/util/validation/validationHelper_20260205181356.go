// This package holds the the logic to validate selected options.
// The values available to choose are defined within the LOWER_LIMIT and UPPER_LIMIT constants in the config package.
package validation

import c "bucles/config"

// Returns true if the selected option value is in between the values defined in LOWER_LIMIT and UPPER_LIMIT
func IsAValidOption(selectedOption int) bool {

	return (c.LOWER_LIMIT >= selectedOption) && (selectedOption <= c.UPPER_LIMIT)

}

func IsExitSelected(selectedOption int) bool {

	return selectedOption == c.UPPER_LIMIT

}
