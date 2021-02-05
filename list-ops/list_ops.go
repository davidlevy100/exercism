// Package listops implements basic list operations.
package listops

type binFunc func(x, y int) int
type predFunc func(x int) bool
type unaryFunc func(x int) int

// IntList is a list of type int
type IntList []int

// Append adds one list to another
func (a IntList) Append(b IntList) []int {

	aLength := len(a)
	bLength := len(b)

	newList := make(IntList, aLength+bLength)

	for i := range newList {

		if i < aLength {
			newList[i] = a[i]
		} else {
			newList[i] = b[i-aLength]
		}

	}

	return newList

}

// Concat flattens a list of lists
func (a IntList) Concat(b []IntList) IntList {

	result := append(IntList{}, a...)

	for _, subList := range b {
		result = append(result, subList...)
	}

	return result
}

// Filter returns the list of all items for which `predicate(item)` is True, given a predicate and a list,
func (a IntList) Filter(thisFunc predFunc) IntList {
	result := IntList{}

	for _, value := range a {
		if thisFunc(value) {
			result = append(result, value)
		}
	}

	return result
}

// Foldl will fold (reduce) each item into the accumulator from the left using `function(accumulator, item)`, given a function, a list, and initial accumulator
func (a IntList) Foldl(thisFunc binFunc, initial int) int {
	for _, value := range a {
		initial = thisFunc(initial, value)
	}

	return initial
}

// Foldr will fold (reduce) each item into the accumulator from the right using `function(accumulator, item)`, given a function, a list, and initial accumulator
func (a IntList) Foldr(thisFunc binFunc, initial int) int {

	for _, value := range a.Reverse() {
		initial = thisFunc(value, initial)
	}

	return initial
}

// Length will return the length of an IntList
func (a IntList) Length() int {
	result := 0

	for range a {
		result++
	}
	return result
}

// Map returns the list of the results of applying `function(item)` on all items given a function and a list.
func (a IntList) Map(thisFunc unaryFunc) IntList {
	result := make(IntList, a.Length())

	for i, value := range a {
		result[i] = thisFunc(value)
	}
	return result
}

// Reverse returns an IntList in reverse order
func (a IntList) Reverse() IntList {

	thisLength := a.Length()

	result := make(IntList, thisLength)

	for i, value := range a {
		result[thisLength-i-1] = value
	}

	return result

}
