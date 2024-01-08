// Package diffsquares offers functions to handle the difference of squares
package diffsquares

// SquareOfSum returns the square of a sum of numbers
func SquareOfSum(n int) int {

	total := (n * (n + 1)) / 2

	return total * total

}

// SumOfSquares returns the sum of a square of numbers
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference returns the difference between sum of squares and the square of sum
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
