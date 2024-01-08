// Package strand provides functions to encode RNA
package strand

import (
	"fmt"
	"strings"
)

var complements = map[string]string{
	"G": "C",
	"C": "G",
	"T": "A",
	"A": "U",
}

// ToRNA encodes an RNA strand from a given DNA striand
func ToRNA(dna string) string {

	var result strings.Builder
	var complement string

	for _, letter := range dna {
		complement = complements[string(letter)]
		fmt.Fprintf(&result, "%s", complement)
	}

	return result.String()

}
