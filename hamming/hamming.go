// Package hamming provides tools to determine the distance between two starnds of DNA
package hamming

import "errors"

// Distance calculates the distance between two strings
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("sequences not equal")
	}

	result := 0
	aRunes := []rune(a)
	bRunes := []rune(b)

	for i := range aRunes {
		if aRunes[i] != bRunes[i] {
			result++
		}
	}

	return result, nil

}
