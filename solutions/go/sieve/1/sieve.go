// Package sieve implements the Sieve of Eratosthenes.
package sieve

// Sieve returns all primes ≤ limit.
func Sieve(limit int) []int {
	if limit < 2 {
		return nil
	}
	p := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		p[i] = true
	}
	for i := 2; i*i <= limit; i++ {
		if p[i] {
			for j := i * i; j <= limit; j += i {
				p[j] = false
			}
		}
	}
	var primes []int
	for i := 2; i <= limit; i++ {
		if p[i] {
			primes = append(primes, i)
		}
	}
	return primes
}
