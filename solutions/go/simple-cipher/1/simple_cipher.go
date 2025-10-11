// Package cipher implements simple substitution ciphers including
// Caesar, Shift, and Vigenère ciphers.
package cipher

import "strings"

type shift struct {
	distance int
}

type vigenere struct {
	key string
}

// NewCaesar returns a Caesar cipher with a fixed shift of 3.
func NewCaesar() Cipher {
	return NewShift(3)
}

// NewShift returns a cipher with a constant shift distance.
// Returns nil if distance is 0, >25, or <-25.
func NewShift(distance int) Cipher {
	if distance == 0 || distance > 25 || distance < -25 {
		return nil
	}
	return shift{distance}
}

// Encode applies a positive shift to all letters in the input.
func (c shift) Encode(s string) string { return c.transform(s, c.distance) }

// Decode applies a negative shift to all letters in the input.
func (c shift) Decode(s string) string { return c.transform(s, -c.distance) }

// transform shifts letters in the input by the given offset.
func (c shift) transform(input string, offset int) string {
	var result strings.Builder
	for _, r := range strings.ToLower(input) {
		if r >= 'a' && r <= 'z' {
			result.WriteRune(shiftRune(r, offset))
		}
	}
	return result.String()
}

// NewVigenere returns a cipher that shifts each letter by
// the position of its corresponding key letter.
// Returns nil for empty, invalid, or all-'a' keys.
func NewVigenere(key string) Cipher {
	if len(key) == 0 {
		return nil
	}

	allA := true
	for _, r := range key {
		if r < 'a' || r > 'z' {
			return nil
		}
		if r != 'a' {
			allA = false
		}
	}
	if allA {
		return nil
	}

	return vigenere{key}
}

// Encode encodes text using the Vigenère cipher.
func (v vigenere) Encode(s string) string { return v.transform(s, 1) }

// Decode decodes text using the Vigenère cipher.
func (v vigenere) Decode(s string) string { return v.transform(s, -1) }

// transform encodes or decodes input depending on the sign.
func (v vigenere) transform(input string, sign int) string {
	var result strings.Builder
	key := v.key
	keyLen := len(key)
	keyIndex := 0

	for _, r := range strings.ToLower(input) {
		if r < 'a' || r > 'z' {
			continue
		}
		shift := int(key[keyIndex%keyLen] - 'a')
		result.WriteRune(shiftRune(r, sign*shift))
		keyIndex++
	}
	return result.String()
}

// shiftRune returns a letter shifted within 'a'–'z' by offset.
func shiftRune(r rune, offset int) rune {
	return 'a' + ((r-'a')+rune(offset)+26)%26
}
