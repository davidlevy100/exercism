// Package bottlesong generates verses of the “Ten Green Bottles” song.
package bottlesong

import (
	"fmt"
	"strings"
)

// Recite returns verses of the Green Bottles song.
func Recite(startBottles, takeDown int) []string {
	words := []string{
		"No", "One", "Two", "Three", "Four",
		"Five", "Six", "Seven", "Eight", "Nine", "Ten",
	}

	bottle := func(n int) string {
		if n == 1 {
			return "bottle"
		}
		return "bottles"
	}

	var verses []string

	for i := startBottles; i > startBottles-takeDown; i-- {
		// main verse
		line := fmt.Sprintf("%s green %s hanging on the wall,", words[i], bottle(i))
		verses = append(verses,
			line,
			line,
			"And if one green bottle should accidentally fall,",
			fmt.Sprintf("There'll be %s green %s hanging on the wall.",
				strings.ToLower(words[i-1]), bottle(i-1)),
		)

		// blank line between verses
		if i > startBottles-takeDown+1 {
			verses = append(verses, "")
		}
	}

	return verses
}
