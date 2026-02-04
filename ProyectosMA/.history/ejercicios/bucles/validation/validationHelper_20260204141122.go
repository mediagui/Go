// This package holds the the logic to validate selected options.
// The values available to choose are defined within the LOWER_LIMIT and UPPER_LIMIT constants.
package validation

const LOWER_LIMIT = 1
const UPPER_LIMIT = 6

// Returns true if the selected option value is in between the values defined in LOWER_LIMIT and UPPER_LIMIT
func IsAValidOption(selectedOption int) bool {

	return (LOWER_LIMIT >= selectedOption) && (selectedOption <= UPPER_LIMIT)

}
