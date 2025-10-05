// Package luhn validates strings of digits using the Luhn algorithm.
package luhn

import (
	"strconv"
	"strings"
)

// Valid reports whether s passes the Luhn checksum test.
func Valid(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	if len(s) <= 1 {
		return false
	}

	var total int
	parity := false // false = no double, true = double

	for i := len(s) - 1; i >= 0; i-- {
		r := s[i]

		if r < '0' || r > '9' {
			return false
		}

		n, _ := strconv.Atoi(string(r))

		if parity {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}

		total += n
		parity = !parity
	}

	return total%10 == 0
}
