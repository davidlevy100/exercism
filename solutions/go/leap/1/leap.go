// Package leap checks if a year is a leap year.
package leap

// IsLeapYear reports if year is a leap year.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
