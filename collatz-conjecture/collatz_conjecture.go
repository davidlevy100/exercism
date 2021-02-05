// Package collatzconjecture calculates the steps of the collatz conjecture
package collatzconjecture

import (
	"errors"
	"math"
)

// CollatzConjecture takes a positive integer input and returns the number of steps needed to get to 1.
func CollatzConjecture(input int) (int, error) {

	if input < 1 {
		return 0, errors.New("Positive integers only, please")
	}

	steps := 0

	n := float64(input)

	for {

		if n == 1.0 {
			break
		}

		if math.Mod(n, 2.0) == 0.0 {
			n = n / 2.0
		} else {
			n = (n * 3) + 1
		}

		steps++
	}

	return steps, nil

}
