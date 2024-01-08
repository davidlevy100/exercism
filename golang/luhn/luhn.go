// Package luhn offers functions to validate strigns of integers
package luhn

import (
	"strconv"
	"strings"
)

// Valid tests whether a string of integers is passes the Luhn formula
func Valid(s string) bool {

	s = strings.ReplaceAll(s, " ", "")

	if len(s) <= 1 {
		return false
	}

	var total int
	var parity bool

	// Process string in reverse
	for i := len(s) - 1; i >= 0; i-- {

		// all spaces have been removed, so any non-digit invalidates the string
		newVal, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return false
		}

		// using FTFTFTFT pattern to check every other digit
		if parity {

			newVal *= 2

			if newVal > 9 {
				newVal -= 9
			}
			parity = false

		} else {
			parity = true
		}

		total += newVal

	}

	return total%10 == 0

}
