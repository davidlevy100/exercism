// Package scrabble provides functions for the scrabble game
package scrabble

// Score tallies the letter scores for a scrabble word
func Score(word string) int {
	var result int

	for _, letter := range word {
		switch letter {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T', 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			result++
		case 'D', 'G', 'd', 'g':
			result += 2
		case 'B', 'C', 'M', 'P', 'b', 'c', 'm', 'p':
			result += 3
		case 'F', 'H', 'V', 'W', 'Y', 'f', 'h', 'v', 'w', 'y':
			result += 4
		case 'K', 'k':
			result += 5
		case 'J', 'X', 'j', 'x':
			result += 8
		case 'Q', 'Z', 'q', 'z':
			result += 10
		}
	}
	return result
}
