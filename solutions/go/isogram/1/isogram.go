// Package isogram checks if a word or phrase is an isogram.
package isogram

import "unicode"

// IsIsogram reports whether a word or phrase has no repeating letters.
// It ignores case and non-letter characters.
func IsIsogram(word string) bool {
	seen := make(map[rune]bool)

	for _, r := range word {
		lr := unicode.ToLower(r)
		if lr < 'a' || lr > 'z' {
			continue
		}
		if seen[lr] {
			return false
		}
		seen[lr] = true
	}
	return true
}
