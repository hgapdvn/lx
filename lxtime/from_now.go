package lxtime

import "time"

// FromNow returns a time that is n units in the future.
// It is equivalent to time.Now().Add(n * unit).
//
// Example:
//
//	fiveMinutesLater := lxtime.FromNow(5, time.Minute)
//	// fiveMinutesLater: approximately 5 minutes from now
//
//	oneHourLater := lxtime.FromNow(1, time.Hour)
//	// oneHourLater: approximately 1 hour from now
func FromNow(n int, unit time.Duration) time.Time {
	return time.Now().Add(time.Duration(n) * unit)
}
