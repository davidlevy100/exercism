// Package parsinglogfiles provides helpers for analyzing custom log formats.
package parsinglogfiles

import "regexp"

// Precompiled regexes for efficiency.
var (
	validLineRe = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	splitRe     = regexp.MustCompile(`<[-~=*]*>`)
	passwordRe  = regexp.MustCompile(`(?i)".*password.*"`)
	endOfLineRe = regexp.MustCompile(`end-of-line\d+`)
	userNameRe  = regexp.MustCompile(`User\s+(\S+)`)
)

// IsValidLine reports whether the line starts with a known log level.
func IsValidLine(text string) bool {
	return validLineRe.MatchString(text)
}

// SplitLogLine splits a line into fields using custom "<...>" separators.
func SplitLogLine(text string) []string {
	return splitRe.Split(text, -1)
}

// CountQuotedPasswords counts lines with "password" (any case) inside quotes.
func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, line := range lines {
		if passwordRe.MatchString(line) {
			count++
		}
	}
	return count
}

// RemoveEndOfLineText strips "end-of-line" artifacts followed by digits.
func RemoveEndOfLineText(text string) string {
	return endOfLineRe.ReplaceAllString(text, "")
}

// TagWithUserName prefixes lines containing "User <name>" with "[USR] <name>".
func TagWithUserName(lines []string) []string {
	out := make([]string, len(lines))
	for i, line := range lines {
		if m := userNameRe.FindStringSubmatch(line); m != nil {
			out[i] = "[USR] " + m[1] + " " + line
		} else {
			out[i] = line
		}
	}
	return out
}
