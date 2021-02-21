// Package grains provides methods for the wheat on a chessboard problem
package grains

import (
	"errors"
)

// Square returns how many grains on a given square, as an integer
func Square(n int) (uint64, error) {

	if n < 1 || n > 64 {
		return 0, errors.New("input out of range")
	}

	// any particular square is 2^n-1
	return 1 << (n - 1), nil

}

// Total returns the total number of grains on the chessboard, as an integer
func Total() uint64 {

	// solution to the Wheat and Chessboard problem is 2^64 - 1
	return 1<<64 - 1

}
