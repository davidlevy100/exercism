// Package isogram provides tests for strings
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns true if no letters are repeated, else false
func IsIsogram(s string) bool {

	result := true

	seenRune := make(map[rune]bool)

	for _, thisRune := range strings.ToLower(s) {
		if unicode.IsLetter(thisRune) {
			seen := seenRune[thisRune]

			if !seen {
				seenRune[thisRune] = true
			} else {
				result = false
				break
			}

		}
	}

	return result

}
