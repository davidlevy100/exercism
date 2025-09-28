// Package booking provides helpers for working with appointment dates and times.
package booking

import (
	"fmt"
	"time"
)

// Schedule returns the parsed time for a date string in "M/D/2006 15:04:05" format.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed reports whether a date in "January 2, 2006 15:04:05" format is in the past.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return time.Now().After(t)
}

// IsAfternoonAppointment reports whether a time in
// "Monday, January 2, 2006 15:04:05" format is 12:00–17:59.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns "You have an appointment on <weekday, month day, year>, at <HH:MM>."
// for a date in "M/D/2006 15:04:05" format.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
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
