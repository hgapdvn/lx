package lxtime

import "time"

// EndOfQuarter returns the end of the quarter for the given time.
// Returns the last day of the quarter at 23:59:59.999999999 in the same timezone.
// Q1: March 31, Q2: June 30, Q3: September 30, Q4: December 31
//
// Example:
//
//	t := time.Date(2026, 6, 15, 15, 30, 45, 0, time.UTC)
//	end := lxtime.EndOfQuarter(t)
//	// end: 2026-06-30 23:59:59.999999999 +0000 UTC (Q2)
func EndOfQuarter(t time.Time) time.Time {
	return StartOfQuarter(t).
		AddDate(0, 3, 0).
		Add(-time.Nanosecond)
}
