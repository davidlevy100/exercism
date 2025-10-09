// Package prime provides a function to find the nth prime.
package prime

import (
	"errors"
	"math"
)

// Nth returns the nth prime number using previously found primes
// to reduce redundant divisibility checks.
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n must be greater than 0")
	}

	primes := []int{}

	for i := 2; ; i++ {
		if isPrime(i, primes) {
			primes = append(primes, i)
			if len(primes) == n {
				return i, nil
			}
		}
	}
}

// isPrime checks divisibility only against known smaller primes up to √x.
func isPrime(x int, primes []int) bool {
	limit := int(math.Sqrt(float64(x)))
	for _, p := range primes {
		if p > limit {
			break
		}
		if x%p == 0 {
			return false
		}
	}
	return true
}
