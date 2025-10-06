// Package bob responds to remarks like a lackadaisical teenager.
package bob

import "strings"

// Hey returns Bob's response to the given remark.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	isQuestion := strings.HasSuffix(remark, "?")
	isYelling := strings.ToUpper(remark) == remark && strings.ContainsAny(remark, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	switch {
	case remark == "":
		return "Fine. Be that way!"
	case isYelling && isQuestion:
		return "Calm down, I know what I'm doing!"
	case isYelling:
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	default:
		return "Whatever."
	}
}
