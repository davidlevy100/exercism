// Package prime provides prime factorization.
package prime

// Factors returns the prime factors of n.
func Factors(n int64) []int64 {
	var factors []int64
	for i := int64(2); i*i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}
