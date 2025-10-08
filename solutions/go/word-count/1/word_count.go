// Package wordcount counts the frequency of words in a phrase.
package wordcount

import (
	"strings"
	"unicode"
)

// Frequency maps words to their occurrence count.
type Frequency map[string]int

// WordCount returns a map of word frequencies in the given phrase.
// It treats apostrophes within words as valid (e.g., "don't"),
// but strips leading or trailing apostrophes.
func WordCount(phrase string) Frequency {
	result := Frequency{}
	var b strings.Builder
	b.Grow(len(phrase))

	// Normalize to lowercase, keeping only letters, digits, and apostrophes.
	for _, r := range phrase {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\'':
			b.WriteRune(unicode.ToLower(r))
		default:
			b.WriteRune(' ')
		}
	}

	// Split on whitespace
	for _, raw := range strings.Fields(b.String()) {
		w := strings.Trim(raw, "'") // remove leading/trailing apostrophes
		if w != "" {
			result[w]++
		}
	}

	return result
}
