// Package rotationalcipher implements a simple rotational cipher (Caesar cipher).
package rotationalcipher

import "unicode"

// RotationalCipher applies a Caesar cipher with the given shiftKey to plain text.
// Non-alphabetic characters are left unchanged.
func RotationalCipher(plain string, shiftKey int) string {
	result := make([]rune, len(plain))

	for i, r := range plain {
		switch {
		case unicode.IsUpper(r):
			result[i] = 'A' + (r-'A'+rune(shiftKey))%26
		case unicode.IsLower(r):
			result[i] = 'a' + (r-'a'+rune(shiftKey))%26
		default:
			result[i] = r
		}
	}

	return string(result)
}
