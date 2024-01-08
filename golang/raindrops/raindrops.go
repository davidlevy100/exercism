// Package raindrops converts numbers to strings with special rules
package raindrops

import (
	"strconv"
)

// Convert returns raindrop sounds if the given int is divisible by 3, 5 or 7
func Convert(n int) string {

	var result string

	if n%3 == 0 {
		result += "Pling"
	}

	if n%5 == 0 {
		result += "Plang"
	}

	if n%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		result = strconv.Itoa(n)
	}

	return result

}
