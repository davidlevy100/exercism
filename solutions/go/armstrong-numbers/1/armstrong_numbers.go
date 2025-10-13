// Package armstrong determines whether numbers are Armstrong (narcissistic) numbers.
package armstrong

// IsNumber reports whether n is an Armstrong number.
// An Armstrong number equals the sum of its digits each raised
// to the power of the number of digits.
// Example: 153 = 1³ + 5³ + 3³.
func IsNumber(n int) bool {
	if n < 0 {
		return false
	}

	// Trivial one-digit case: always Armstrong.
	if n < 10 {
		return true
	}

	value := n
	digits := []int{}
	for value > 0 {
		digits = append(digits, value%10)
		value /= 10
	}

	sum := 0
	for _, d := range digits {
		exp := 1
		for range len(digits) { // valid in Go 1.22+
			exp *= d
		}
		sum += exp
	}

	return sum == n
}
