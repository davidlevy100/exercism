// Package accumulate returns a new collection containing the result of applying an operation to each element of the input collection
package accumulate

// Accumulate will perform the given function on the given collection
func Accumulate(sa []string, f func(string) string) []string {

	var result []string

	for _, thisString := range sa {
		result = append(result, f(thisString))
	}

	return result

}
