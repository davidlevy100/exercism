// Package proverb creates proverbs from string arrays.
package proverb

import "fmt"

// Proverb returns the proverb as a slice of strings.
// Given a list of words, it creates the "for want of" chain
// ending with "And all for the want of a X."
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return nil
	}

	result := make([]string, 0, len(rhyme))
	for i := 0; i < len(rhyme)-1; i++ {
		result = append(result,
			fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]),
		)
	}
	result = append(result, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return result
}
