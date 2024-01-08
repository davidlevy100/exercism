// Package proverb creates proverbs from string arrays
package proverb

import "fmt"

// Proverb returns a proverb as an array of strings
func Proverb(rhyme []string) []string {

	var result []string

	for i := len(rhyme) - 1; i >= 0; i-- {
		if i == 0 {
			result = append(result, fmt.Sprintf("And all for the want of a %s.", rhyme[i]))
		} else {
			s := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i])
			result = append([]string{s}, result...)
		}
	}

	return result
}
