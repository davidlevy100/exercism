// Package stringset implements a basic string set type with standard set operations.
package stringset

import (
	"fmt"
	"strings"
)

// Set represents a collection of unique string values.
type Set map[string]struct{}

// New returns an empty Set.
func New() Set {
	return Set{}
}

// NewFromSlice returns a new Set initialized with the given slice of strings.
func NewFromSlice(l []string) Set {
	result := New()
	for _, s := range l {
		result[s] = struct{}{}
	}
	return result
}

// String returns a formatted string representation of the set.
// Example: {"a", "b"}. The order of elements is not guaranteed.
func (s Set) String() string {
	var sb strings.Builder
	sb.WriteString("{")

	i := 0
	for k := range s {
		if i > 0 {
			sb.WriteString(", ")
		}
		fmt.Fprintf(&sb, "%q", k)
		i++
	}

	sb.WriteString("}")
	return sb.String()
}

// IsEmpty reports whether the set contains no elements.
func (s Set) IsEmpty() bool {
	return len(s) == 0
}

// Has reports whether the given element exists in the set.
func (s Set) Has(elem string) bool {
	_, ok := s[elem]
	return ok
}

// Add inserts an element into the set.
func (s Set) Add(elem string) {
	s[elem] = struct{}{}
}

// Subset reports whether s1 is a subset of s2.
func Subset(s1, s2 Set) bool {
	if len(s1) > len(s2) {
		return false
	}
	for k := range s1 {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

// Disjoint reports whether two sets have no elements in common.
func Disjoint(s1, s2 Set) bool {
	return len(Intersection(s1, s2)) == 0
}

// Equal reports whether two sets contain exactly the same elements.
func Equal(s1, s2 Set) bool {
	return len(s1) == len(s2) && Subset(s1, s2)
}

// Intersection returns a new set containing elements common to both sets.
func Intersection(s1, s2 Set) Set {
	result := New()
	for k := range s1 {
		if s2.Has(k) {
			result.Add(k)
		}
	}
	return result
}

// Difference returns a new set containing elements in s1 but not in s2.
func Difference(s1, s2 Set) Set {
	result := New()
	for k := range s1 {
		if !s2.Has(k) {
			result.Add(k)
		}
	}
	return result
}

// Union returns a new set containing all elements from both sets.
func Union(s1, s2 Set) Set {
	result := New()
	for k := range s1 {
		result.Add(k)
	}
	for k := range s2 {
		result.Add(k)
	}
	return result
}
