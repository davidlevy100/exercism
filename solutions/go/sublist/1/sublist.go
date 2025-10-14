// Package sublist provides functionality to compare two integer slices
// and determine whether one is a sublist, superlist, equal to, or unequal
// to the other.
package sublist

import "slices"

// Sublist compares two lists and returns their relation.
func Sublist(l1, l2 []int) Relation {
	switch {
	case len(l1) == 0 && len(l2) == 0:
		return RelationEqual
	case len(l1) == 0:
		return RelationSublist
	case len(l2) == 0:
		return RelationSuperlist
	case slices.Equal(l1, l2):
		return RelationEqual
	}

	var shorter, longer []int
	var relation Relation

	if len(l1) > len(l2) {
		shorter, longer = l2, l1
		relation = RelationSuperlist
	} else {
		shorter, longer = l1, l2
		relation = RelationSublist
	}

	for i := 0; i <= len(longer)-len(shorter); i++ {
		if slices.Equal(shorter, longer[i:i+len(shorter)]) {
			return relation
		}
	}

	return RelationUnequal
}
