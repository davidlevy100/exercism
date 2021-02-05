/*
Package gigasecond adds 10^9 seconds to a given time and returns the new time
*/
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond takes a time.Time type variable and adds 10^9 seconds to it and returns the new time.
func AddGigasecond(t time.Time) time.Time {

	return t.Add(time.Second * 1000000000)

}
