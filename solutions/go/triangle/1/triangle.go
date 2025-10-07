// Package triangle classifies triangles by side length.
package triangle

import "math"

// Kind represents the type of triangle.
type Kind string

// Triangle kinds.
const (
	NaT Kind = "NaT" // not a triangle
	Equ Kind = "Equ" // equilateral
	Iso Kind = "Iso" // isosceles
	Sca Kind = "Sca" // scalene
)

// KindFromSides returns the triangle type for sides a, b, and c.
func KindFromSides(a, b, c float64) Kind {
	switch {
	case notATriangle(a, b, c):
		return NaT
	case equilateral(a, b, c):
		return Equ
	case isosceles(a, b, c):
		return Iso
	default:
		return Sca
	}
}

// notATriangle reports whether sides a, b, and c cannot form a valid triangle.
func notATriangle(a, b, c float64) bool {
	switch {
	case a <= 0 || b <= 0 || c <= 0:
		return true
	case math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c):
		return true
	case math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0):
		return true
	case a+b < c || a+c < b || b+c < a:
		return true
	default:
		return false
	}
}

// equilateral reports whether all sides are equal.
func equilateral(a, b, c float64) bool {
	return a == b && b == c
}

// isosceles reports whether at least two sides are equal.
func isosceles(a, b, c float64) bool {
	return a == b || b == c || a == c
}
