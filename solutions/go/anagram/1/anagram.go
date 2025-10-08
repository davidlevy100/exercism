// Package anagram detects anagrams in a list of words.
package anagram

import "strings"

// Detect returns the words from candidates that are anagrams of subject.
func Detect(subject string, candidates []string) []string {
	result := make([]string, 0, len(candidates))
	s := strings.ToLower(subject)

	for _, c := range candidates {
		if isAnagram(s, strings.ToLower(c)) {
			result = append(result, c)
		}
	}

	return result
}

// isAnagram reports whether two lowercased words are anagrams.
func isAnagram(a, b string) bool {
	if a == b || len(a) != len(b) {
		return false
	}

	counts := make(map[rune]int, len(a))

	for _, r := range a {
		counts[r]++
	}
	for _, r := range b {
		counts[r]--
	}

	for _, v := range counts {
		if v != 0 {
			return false
		}
	}

	return true
}
