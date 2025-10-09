// Package primes provides a function to find the nth prime number.
package prime

import (
	"errors"
	"math"
)

// Nth returns the nth prime number using a Sieve of Eratosthenes.
// It estimates an upper bound using the prime number theorem for efficiency.
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n must be >= 1")
	}
	if n == 1 {
		return 2, nil
	}

	// Estimate an upper bound for the nth prime (pₙ ≈ n(log n + log log n))
	fn := float64(n)
	limit := int(fn*(math.Log(fn)+math.Log(math.Log(fn)))) + 10 // small buffer

	sieve := make([]bool, limit+1)
	primes := []int{}

	for i := 2; i <= limit; i++ {
		if !sieve[i] {
			primes = append(primes, i)
			for j := i * i; j <= limit; j += i {
				sieve[j] = true
			}
			if len(primes) == n {
				return i, nil
			}
		}
	}

	// Should never happen unless estimate was too low
	return 0, errors.New("limit too low to find nth prime")
}
