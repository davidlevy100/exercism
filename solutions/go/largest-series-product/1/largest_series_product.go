// Package lsproduct computes the largest product of consecutive digits in a string.
package lsproduct

import "errors"

// LargestSeriesProduct returns the largest product of any span-length substring.
// Span must be non-negative, not longer than digits, and digits must be numeric.
// A span of 0 returns 1 (empty product).
func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 {
		return 0, errors.New("span must be non-negative")
	}
	if span > len(digits) {
		return 0, errors.New("span longer than string length")
	}
	if span == 0 {
		return 1, nil
	}

	var max int64
	for i := 0; i+span <= len(digits); i++ {
		var product int64 = 1
		for j := i; j < i+span; j++ {
			d := digits[j]
			if d < '0' || d > '9' {
				return 0, errors.New("digits must be numeric")
			}
			product *= int64(d - '0')
		}
		if product > max {
			max = product
		}
	}
	return max, nil
}
