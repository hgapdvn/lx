package lxtime

import "time"

// Ago returns a time that is n units in the past.
// It is equivalent to time.Now().Add(-n * unit).
//
// Example:
//
//	fiveMinutesAgo := lxtime.Ago(5, time.Minute)
//	// fiveMinutesAgo: approximately 5 minutes before now
//
//	oneHourAgo := lxtime.Ago(1, time.Hour)
//	// oneHourAgo: approximately 1 hour before now
func Ago(n int, unit time.Duration) time.Time {
	return time.Now().Add(-time.Duration(n) * unit)
}
