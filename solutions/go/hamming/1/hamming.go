// Package hamming provides functions for computing Hamming distance.
package hamming

import "errors"

// ErrUnequalLength is returned when two strings of unequal length are compared.
var ErrUnequalLength = errors.New("hamming: strings must be of equal length")

// Distance computes the Hamming distance between two equal-length strings.
// Returns an error if the input strings differ in length.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, ErrUnequalLength
	}

	diffs := 0
	for i := range a {
		if a[i] != b[i] {
			diffs++
		}
	}
	return diffs, nil
}
