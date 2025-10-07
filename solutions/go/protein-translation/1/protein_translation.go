// Package protein translates RNA sequences into proteins.
package protein

import "errors"

// Sentinel errors for invalid or special codons.
var (
	ErrInvalidBase = errors.New("invalid base")
	ErrStop        = errors.New("stop codon found")
)

// FromRNA translates a full RNA string into a slice of proteins.
// Stops processing when a STOP codon is encountered.
func FromRNA(rna string) ([]string, error) {
	var proteins []string

	for i := 0; i <= len(rna)-3; i += 3 {
		codon := rna[i : i+3]
		aa, err := FromCodon(codon)

		switch err {
		case nil:
			proteins = append(proteins, aa)
		case ErrStop:
			return proteins, nil // normal termination
		default:
			return nil, err // invalid base
		}
	}

	return proteins, nil
}

// FromCodon returns the amino acid for a codon, or ErrStop if it's a stop codon.
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
