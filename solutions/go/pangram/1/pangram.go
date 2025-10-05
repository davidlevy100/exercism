// Package pangram checks if a string contains every letter of the alphabet.
package pangram

import "unicode"

// IsPangram returns true if input includes all 26 letters of the English alphabet.
func IsPangram(input string) bool {
	var seen [26]bool
	count := 0

	for _, r := range input {
		lr := unicode.ToLower(r)
		if lr < 'a' || lr > 'z' {
			continue
		}
		idx := lr - 'a'
		if !seen[idx] {
			seen[idx] = true
			count++
			if count == 26 {
				return true
			}
		}
	}
	return false
}
