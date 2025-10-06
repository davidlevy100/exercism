// Package isbn validates ISBN-10 numbers.
package isbn

import "unicode"

// IsValidISBN reports whether s is a valid ISBN-10.
// Dashes are ignored; 'X' is allowed only as the final digit (value 10).
func IsValidISBN(s string) bool {
	score, multiplier := 0, 10

	for _, r := range s {
		switch {
		case r == '-':
			continue
		case unicode.IsDigit(r):
			score += int(r-'0') * multiplier
			multiplier--
		case r == 'X' && multiplier == 1:
			score += 10
			multiplier--
		default:
			return false
		}
		if multiplier < 0 {
			return false
		}
	}

	return multiplier == 0 && score%11 == 0
}
