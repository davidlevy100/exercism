// Package leap tests for leap years
package leap

// IsLeapYear will return true if the given year is a leap year, else false
func IsLeapYear(year int) bool {

	if year%4 != 0 {
		return false
	}

	if year%100 == 0 {
		if year%400 == 0 {
			return true
		}
		return false
	}

	return true

}
