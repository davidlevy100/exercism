// Package summultiples finds the sum of all unique multiples of given divisors below a limit.
package summultiples

// SumMultiples returns the sum of all unique multiples of the provided divisors below limit.
func SumMultiples(limit int, divisors ...int) int {
	set := make(map[int]bool)
	sum := 0

	for _, d := range divisors {
		if d == 0 {
			continue
		}
		for n := d; n < limit; n += d {
			if !set[n] {
				set[n] = true
				sum += n
			}
		}
	}
	return sum
}
