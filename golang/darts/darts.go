// Package darts has functions for the game of Darts
package darts

import "math"

// Score calculates the score given an x, y position
func Score(x, y float64) int {

	distance := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))

	var result int

	switch {
	case distance <= 1:
		result = 10
	case distance <= 5:
		result = 5
	case distance <= 10:
		result = 1
	default:
		result = 0
	}

	return result

}
