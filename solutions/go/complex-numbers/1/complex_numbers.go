// Package complexnumbers implements operations for complex numbers.
package complexnumbers

import "math"

// Number represents a complex number with real (a) and imaginary (b) parts.
type Number struct {
	a, b float64
}

// Real returns the real part of the number.
func (n Number) Real() float64 { return n.a }

// Imaginary returns the imaginary part of the number.
func (n Number) Imaginary() float64 { return n.b }

// Add returns the sum of n1 and n2.
func (n1 Number) Add(n2 Number) Number {
	return Number{
		a: n1.a + n2.a,
		b: n1.b + n2.b,
	}
}

// Subtract returns the result of n1 - n2.
func (n1 Number) Subtract(n2 Number) Number {
	return Number{
		a: n1.a - n2.a,
		b: n1.b - n2.b,
	}
}

// Multiply returns the product of n1 and n2.
func (n1 Number) Multiply(n2 Number) Number {
	return Number{
		a: n1.a*n2.a - n1.b*n2.b,
		b: n1.a*n2.b + n1.b*n2.a,
	}
}

// Times scales n by a real factor.
func (n Number) Times(factor float64) Number {
	return Number{
		a: n.a * factor,
		b: n.b * factor,
	}
}

// Divide returns the quotient n1 / n2.
func (n1 Number) Divide(n2 Number) Number {
	denom := n2.a*n2.a + n2.b*n2.b
	return Number{
		a: (n1.a*n2.a + n1.b*n2.b) / denom,
		b: (n1.b*n2.a - n1.a*n2.b) / denom,
	}
}

// Conjugate returns the complex conjugate of n.
func (n Number) Conjugate() Number {
	return Number{
		a: n.a,
		b: -n.b,
	}
}

// Abs returns the magnitude (absolute value) of n.
func (n Number) Abs() float64 { return math.Hypot(n.a, n.b) }

// Exp returns e raised to the power of n.
func (n Number) Exp() Number {
	expA := math.Exp(n.a)
	return Number{
		a: expA * math.Cos(n.b),
		b: expA * math.Sin(n.b),
	}
}
