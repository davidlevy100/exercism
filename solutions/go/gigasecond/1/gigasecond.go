// Package gigasecond provides functionality for time calculations.
package gigasecond

import "time"

// AddGigasecond adds 1 gigasecond (10^9 seconds) to the given time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
