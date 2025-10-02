// Package collatzconjecture provides a function to compute Collatz steps.
package collatzconjecture

import "errors"

// CollatzConjecture returns the number of steps to reach 1 using the Collatz rules.
// Returns an error if n is not positive.
func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n must be positive")
	}

	steps := 0
	for n > 1 {
		steps++
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}
	return steps, nil
}
