// Package strand provides functionality for working with DNA strands.
package strand

import "strings"

// ToRNA converts a DNA string into its corresponding RNA string.
func ToRNA(dna string) string {
	var b strings.Builder
	b.Grow(len(dna))

	for _, n := range dna {
		switch n {
		case 'G':
			b.WriteRune('C')
		case 'C':
			b.WriteRune('G')
		case 'T':
			b.WriteRune('A')
		case 'A':
			b.WriteRune('U')
		}
	}

	return b.String()
}
