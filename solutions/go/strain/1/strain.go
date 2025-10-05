// Package strain implements Keep and Discard operations on generic collections.
package strain

// Keep returns a new slice containing all elements for which f returns true.
func Keep[T any](input []T, f func(T) bool) []T {
	var out []T
	for _, v := range input {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}

// Discard returns a new slice containing all elements for which f returns false.
func Discard[T any](input []T, f func(T) bool) []T {
	var out []T
	for _, v := range input {
		if !f(v) {
			out = append(out, v)
		}
	}
	return out
}
