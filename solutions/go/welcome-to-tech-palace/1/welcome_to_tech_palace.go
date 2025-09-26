// Package techpalace provides simple helpers for marketing messages.
package techpalace

import (
	"strings"
)

// WelcomeMessage returns a welcome message with the customer's name in uppercase.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder surrounds the message with lines of '*' of the given length.
func AddBorder(msg string, numStars int) string {
	return strings.Repeat("*", numStars) + "\n" + msg + "\n" + strings.Repeat("*", numStars)
}

// CleanupMessage removes '*' and surrounding whitespace from a message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}
