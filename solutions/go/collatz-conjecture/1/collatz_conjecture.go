// Package collatzconjecture calculates the steps of the collatz conjecture
package collatzconjecture

import "errors"

// CollatzConjecture takes a positive integer input and returns the number of steps needed to get to 1.
func CollatzConjecture(input int) (int, error) {
	if input < 1 {
		return 0, errors.New("only positive integers are allowed")
	}
	steps := 0
	for n := input; n != 1; steps++ {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
	}
	return steps, nil
}