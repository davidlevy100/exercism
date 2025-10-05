// Package scrabble provides functions to compute Scrabble word scores.
package scrabble

import "unicode"

// letterScore maps Scrabble tiles to their point values.
var letterScore = map[rune]int{
	'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1, 'L': 1, 'N': 1, 'R': 1, 'S': 1, 'T': 1,
	'D': 2, 'G': 2,
	'B': 3, 'C': 3, 'M': 3, 'P': 3,
	'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4,
	'K': 5,
	'J': 8, 'X': 8,
	'Q': 10, 'Z': 10,
}

// Score returns the Scrabble score of the given word.
// It ignores non-letter characters and treats lowercase as uppercase.
func Score(word string) int {
	sum := 0
	for _, r := range word {
		if v, ok := letterScore[unicode.ToUpper(r)]; ok {
			sum += v
		}
	}
	return sum
}
