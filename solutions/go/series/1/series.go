// Package series provides utilities for working with string slices.
package series

// All returns all contiguous substrings of length n from s.
// If n is greater than len(s), it returns an empty slice.
func All(n int, s string) []string {
	if n <= 0 || n > len(s) {
		return nil
	}
	result := make([]string, 0, len(s)-n+1)
	for i := n; i <= len(s); i++ {
		result = append(result, s[i-n:i])
	}
	return result
}

// UnsafeFirst returns the first n characters of s without bounds checking.
// It panics if n > len(s).
func UnsafeFirst(n int, s string) string {
	return s[:n]
}

// First safely returns the first n characters of s and a boolean indicating success.
func First(n int, s string) (string, bool) {
	if n > len(s) || n <= 0 {
		return "", false
	}
	return s[:n], true
}
