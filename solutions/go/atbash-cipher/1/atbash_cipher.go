// Package atbash implements the classic Atbash substitution cipher.
package atbash

import (
	"strings"
	"unicode"
)

// Atbash encodes text with the Atbash cipher, preserving digits and grouping in 5s.
func Atbash(s string) string {
	var b strings.Builder
	count := 0

	for _, r := range strings.ToLower(s) {
		switch {
		case unicode.IsLetter(r):
			r = 'z' - (r - 'a')
		case unicode.IsDigit(r):
		default:
			continue
		}

		if count > 0 && count%5 == 0 {
			b.WriteRune(' ')
		}

		b.WriteRune(r)
		count++
	}

	return b.String()
}
