// Package bob reposnds to text input
package bob

import "strings"

// Hey will return a string based on the input string
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)

	switch {
	// Yelling a question
	case (strings.ToUpper(remark) == remark) && strings.ContainsAny(remark, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") && strings.HasSuffix(remark, "?"):
		return "Calm down, I know what I'm doing!"

	// Just Yelling
	case strings.ToUpper(remark) == remark && strings.ContainsAny(remark, "ABCDEFGHIJKLMNOPQRSTUVWXYZ"):
		return "Whoa, chill out!"

	// Just a question
	case strings.HasSuffix(remark, "?"):
		return "Sure."

	// No text
	case len(remark) == 0:
		return "Fine. Be that way!"

	// response to anything else
	default:
		return "Whatever."

	}
}
