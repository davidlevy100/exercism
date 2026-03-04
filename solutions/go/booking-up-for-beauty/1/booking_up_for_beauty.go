// Package booking provides helpers for working with appointment dates and times.
package booking

import (
	"fmt"
	"time"
)

// Schedule returns the parsed time for a date string in "M/D/2006 15:04:05" format.
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return t
}

// HasPassed reports whether a date in "January 2, 2006 15:04:05" format is in the past.
func HasPassed(date string) bool {
	t, _ := time.Parse("January 2, 2006 15:04:05", date)
	return time.Now().After(t)
}

// IsAfternoonAppointment reports whether a time in
// "Monday, January 2, 2006 15:04:05" format is 12:00–17:59.
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	hour := t.Hour()
	return hour >= 12 && hour < 18
}

// Description returns "You have an appointment on <weekday, month day, year>, at <HH:MM>."
func Description(date string) string {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return fmt.Sprintf(
		"You have an appointment on %s, at %s.",
		t.Format("Monday, January 2, 2006"),
		t.Format("15:04"),
	)
}

// AnniversaryDate returns September 15 of the current year at midnight UTC.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}