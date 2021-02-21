// Package strain implements the `keep` and `discard` operation on collections.
package strain

// Ints a slice of integers
type Ints []int

// Lists a matrix of integers
type Lists [][]int

// Strings a slice of strings
type Strings []string

// Keep (Ints) returns a new collection containing those elements where the predicate is true
func (si Ints) Keep(f func(int) bool) Ints {

	var result Ints

	for _, thisItem := range si {
		if f(thisItem) {
			result = append(result, thisItem)
		}
	}

	return result

}

// Discard (Ints) returns a new collection containing those elements where the predicate is false
func (si Ints) Discard(f func(int) bool) Ints {

	var result Ints

	for _, thisItem := range si {
		if !f(thisItem) {
			result = append(result, thisItem)
		}
	}

	return result

}

// Keep (Lists) returns a new collection containing those elements where the predicate is true
func (ml Lists) Keep(f func([]int) bool) Lists {

	var result Lists

	for _, thisItem := range ml {
		if f(thisItem) {
			result = append(result, thisItem)
		}
	}

	return result

}

// Keep (Strings) returns a new collection containing those elements where the predicate is true
func (ss Strings) Keep(f func(string) bool) Strings {

	var result Strings

	for _, thisItem := range ss {
		if f(thisItem) {
			result = append(result, thisItem)
		}
	}

	return result

}
