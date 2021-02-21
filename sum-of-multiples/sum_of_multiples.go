// Package summultiples finds the sum of all the unique multiples of particular numbers up to but not including that number
package summultiples

// SumMultiples returns the sum of all the unique multiples of particular numbers up to but not including the limit
func SumMultiples(limit int, divisors ...int) int {

	var result int

	set := make(map[int]bool)

	for _, thisNum := range divisors {

		if thisNum == 0 {
			continue
		}

		for i := thisNum; i < limit; i += thisNum {
			if seen := set[i]; !seen {
				set[i] = true
				result += i
			}
		}
	}

	return result

}
