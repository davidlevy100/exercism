// Package triangle determines the type of triangle
package triangle

import "math"

// Kind is the possible type of triangle
type Kind string

const (
	// NaT Not a Triangle
	NaT Kind = "Nat"
	// Equ Equilateral
	Equ Kind = "Equ"
	// Iso Icsosoles
	Iso Kind = "Iso"
	// Sca Scalene
	Sca Kind = "Sca"
)

// KindFromSides determines the type of triangle by the side lengths
func KindFromSides(a, b, c float64) Kind {

	var k Kind

	if notATriangle(a, b, c) {
		k = NaT
	} else if equilateral(a, b, c) {
		k = Equ
	} else if isosceles(a, b, c) {
		k = Iso
	} else {
		k = Sca
	}

	return k

}

func notATriangle(a, b, c float64) bool {

	result := false

	//any side zero or less
	if a <= 0 || b <= 0 || c <= 0 {
		result = true
	} else if a+b < c || a+c < b || b+c < a {
		result = true
	} else if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		result = true
	} else if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		result = true
	}

	return result

}

func equilateral(a, b, c float64) bool {

	result := false

	if a == b && b == c {
		result = true
	}

	return result
}

func isosceles(a, b, c float64) bool {

	result := false

	if a == b || b == c || a == c {
		result = true
	}

	return result
}
