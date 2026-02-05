// This package holds the the logic to validate selected options.
// The values available to choose are defined within the LOWER_LIMIT and UPPER_LIMIT constants in the config package.
package validation

import "bucles/config"

// Returns true if the selected option value is in between the values defined in LOWER_LIMIT and UPPER_LIMIT
func IsAValidOption(selectedOption int) bool {

	return (config.LOWER_LIMIT >= selectedOption) && (selectedOption <= config.UPPER_LIMIT)

}

func ShouldIExit() {
	if selectedOption == config.UPPER_LIMIT {

		break
	}
}
