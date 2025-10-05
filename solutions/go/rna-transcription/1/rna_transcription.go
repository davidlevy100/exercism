// Package strand provides functions to transcribe RNA from DNA.
package strand

import "strings"

var complements = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA returns the RNA complement of the given DNA strand.
func ToRNA(dna string) string {
	var b strings.Builder
	b.Grow(len(dna))

	for _, r := range dna {
		if comp, ok := complements[r]; ok {
			b.WriteRune(comp)
		}
	}

	return b.String()
}
