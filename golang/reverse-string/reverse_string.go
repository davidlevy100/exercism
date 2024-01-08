// Package reverse offers functions to reverse inputs
package reverse

import "strings"

// Reverse returns a string in reverse
func Reverse(s string) string {

	var sa []string

	for _, letter := range s {
		sa = append([]string{string(letter)}, sa...)
	}

	return strings.Join(sa, "")
}
