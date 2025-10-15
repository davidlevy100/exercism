// Package listops implements basic list operations.
package listops

type binFunc func(x, y int) int
type predFunc func(x int) bool
type unaryFunc func(x int) int

// IntList is a list of integers.
type IntList []int

// Append adds all elements of list b to the end of list a.
func (a IntList) Append(b IntList) IntList {
	newList := make(IntList, len(a)+len(b))
	for i := 0; i < len(a); i++ {
		newList[i] = a[i]
	}
	for i := 0; i < len(b); i++ {
		newList[len(a)+i] = b[i]
	}
	return newList
}

// Concat flattens a list of lists into a single list.
func (a IntList) Concat(lists []IntList) IntList {
	result := a
	for _, list := range lists {
		result = result.Append(list)
	}
	return result
}

// Filter returns all elements that satisfy the given predicate.
func (a IntList) Filter(f predFunc) IntList {
	result := IntList{}
	for _, v := range a {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// Length returns the number of elements in the list.
func (a IntList) Length() int {
	count := 0
	for range a {
		count++
	}
	return count
}

// Map applies a function to every element in the list.
func (a IntList) Map(f unaryFunc) IntList {
	result := make(IntList, len(a))
	for i, v := range a {
		result[i] = f(v)
	}
	return result
}

// Foldl folds the list from the left with the given binary function.
func (a IntList) Foldl(f binFunc, initial int) int {
	acc := initial
	for _, v := range a {
		acc = f(acc, v)
	}
	return acc
}

// Foldr folds the list from the right with the given binary function.
func (a IntList) Foldr(f binFunc, initial int) int {
	acc := initial
	for i := len(a) - 1; i >= 0; i-- {
		acc = f(a[i], acc)
	}
	return acc
}

// Reverse returns a new list in reverse order.
func (a IntList) Reverse() IntList {
	n := len(a)
	result := make(IntList, n)
	for i, v := range a {
		result[n-1-i] = v
	}
	return result
}
