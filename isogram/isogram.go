// Package isogram provides tests for strings
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns true if no letters are repeated, else false
func IsIsogram(s string) bool {

	result := true

	runeCount := make(map[rune]int)

	for _, thisRune := range strings.ToLower(s) {
		if unicode.IsLetter(thisRune) {
			_, ok := runeCount[thisRune]
			if !ok {
				runeCount[thisRune] = 1
			} else {
				result = false
				break
			}
		}
	}

	return result

}
