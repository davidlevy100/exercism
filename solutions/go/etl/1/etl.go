// Package etl provides functionality to transform legacy scoring data
// into a new format with lowercase keys and integer values.
package etl

import "strings"

// Transform converts a map of integer keys to string slices
// into a map of lowercase strings to their corresponding integer.
func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int)
	for k, vals := range in {
		for _, s := range vals {
			out[strings.ToLower(s)] = k
		}
	}
	return out
}
