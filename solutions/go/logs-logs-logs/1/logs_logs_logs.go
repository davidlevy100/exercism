// Package logs provides utilities for classifying, modifying, and validating log messages.
package logs

// Application returns the app name for the first recognized symbol in log;
// otherwise it returns "default".
func Application(log string) string {
	for _, r := range log {
		switch r {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
	}
	return "default"
}

// Replace returns a copy of log with all oldRune runes replaced by newRune.
func Replace(log string, oldRune, newRune rune) string {
	runes := []rune(log)
	for i, r := range runes {
		if r == oldRune {
			runes[i] = newRune
		}
	}
	return string(runes)
}

// WithinLimit reports whether log contains limit or fewer runes.
func WithinLimit(log string, limit int) bool {
	return len([]rune(log)) <= limit
}
