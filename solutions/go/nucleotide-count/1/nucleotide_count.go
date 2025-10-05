// Package dna provides utilities for counting nucleotides.
package dna

import "errors"

// Histogram maps nucleotides to their occurrence counts.
type Histogram map[rune]int

// DNA represents a strand of DNA as a string of nucleotides.
type DNA string

// Counts returns a histogram of valid nucleotides in the DNA strand.
// It returns an error if any invalid nucleotide is present.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	for _, r := range d {
		if _, ok := h[r]; !ok {
			return h, errors.New("invalid nucleotide found")
		}
		h[r]++
	}
	return h, nil
}
