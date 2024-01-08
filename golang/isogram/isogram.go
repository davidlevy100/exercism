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

			if seenRune[thisRune] {
				result = false
				break
			}

			seenRune[thisRune] = true

		}
	}

	return result

}
