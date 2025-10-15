// Package acronym creates acronyms from strings.
package acronym

import "strings"

// Abbreviate converts a phrase into its uppercase acronym.
func Abbreviate(s string) string {
	isSeparator := func(r rune) bool {
		return r == ' ' || r == '-' || r == '_'
	}

	words := strings.FieldsFunc(s, isSeparator)
	var b strings.Builder
	for _, w := range words {
		if len(w) > 0 {
			b.WriteRune([]rune(w)[0])
		}
	}
	return strings.ToUpper(b.String())
}
