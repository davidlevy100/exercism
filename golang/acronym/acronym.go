// Package acronym creates acronyms from strings
package acronym

import (
	"strings"
)

// Abbreviate converts a string input to an acronym as a string.
func Abbreviate(s string) string {

	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")

	words := strings.Split(s, " ")

	var str strings.Builder

	for _, word := range words {

		word = strings.TrimSpace(word)

		if len(word) > 0 {
			str.WriteString(string(word[:1]))
		}

	}

	return strings.ToUpper(str.String())

}
