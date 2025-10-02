// Package raindrops provides a simple converter that maps numbers
// to sounds based on divisibility rules.
package raindrops

import "strconv"

// Convert returns a string according to the raindrop rules:
//   - If divisible by 3, include "Pling".
//   - If divisible by 5, include "Plang".
//   - If divisible by 7, include "Plong".
//   - Otherwise return the number itself as a string.
func Convert(number int) string {
	result := ""

	if number%3 == 0 {
		result += "Pling"
	}
	if number%5 == 0 {
		result += "Plang"
	}
	if number%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		return strconv.Itoa(number)
	}
	return result
}
