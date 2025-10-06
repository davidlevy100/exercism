// Package grains computes rice counts on a chessboard using powers of two.
package grains

import "fmt"

// Square returns grains on square n (1–64), doubling each time from 1.
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, fmt.Errorf("square out of range: %d", n)
	}
	return 1 << (n - 1), nil
}

// Total returns total grains on all 64 squares.
func Total() uint64 {
	return 1<<64 - 1
}
