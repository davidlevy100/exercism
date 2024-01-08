// Package etl provides methods to Extract, Transform and Load data
package etl

import "strings"

// Transform (map[int][]string) takes a letters per score map and returns a score per letter map
func Transform(oldMap map[int][]string) map[string]int {

	result := make(map[string]int)

	for key, value := range oldMap {
		for _, letter := range value {
			result[strings.ToLower(letter)] = key
		}
	}
	return result
}
