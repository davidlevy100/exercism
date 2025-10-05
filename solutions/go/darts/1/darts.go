// Package darts computes the dartboard score based on x,y coordinates.
package darts

import "math"

// Score returns the score for a dart landing at position (x, y).
func Score(x, y float64) int {
	distance := math.Hypot(x, y) // cleaner than Sqrt(Pow())

	switch {
	case distance <= 1:
		return 10
	case distance <= 5:
		return 5
	case distance <= 10:
		return 1
	default:
		return 0
	}
}
