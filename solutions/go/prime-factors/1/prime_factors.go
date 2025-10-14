// Package prime provides prime factorization.
package prime

// Factors returns the prime factors of n.
func Factors(n int64) []int64 {
	result := []int64{}
	for i := int64(2); i*i <= n; i++ {
		for n%i == 0 {
			result = append(result, i)
			n /= i
		}
	}
	if n > 1 {
		result = append(result, n)
	}
	return result
}
