// Package clock implements a simple 24-hour clock that supports
// addition and subtraction of minutes without tracking dates.
package clock

import "fmt"

// Clock represents a time of day without a date.
type Clock struct {
	mins int
}

// normalize keeps minutes in the range [0, 1440).
func normalize(m int) int {
	m %= 24 * 60
	if m < 0 {
		m += 24 * 60
	}
	return m
}

// New returns a new Clock set to h hours and m minutes.
func New(h, m int) Clock {
	return Clock{mins: normalize(h*60 + m)}
}

// Add returns a new Clock advanced by m minutes.
func (c Clock) Add(m int) Clock {
	return Clock{mins: normalize(c.mins + m)}
}

// Subtract returns a new Clock moved backward by m minutes.
func (c Clock) Subtract(m int) Clock {
	return Clock{mins: normalize(c.mins - m)}
}

// String returns the clock time in HH:MM format.
func (c Clock) String() string {
	h := c.mins / 60
	m := c.mins % 60
	return fmt.Sprintf("%02d:%02d", h, m)
}
