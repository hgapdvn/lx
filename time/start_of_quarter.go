package lxtime

import "time"

// StartOfQuarter returns the start of the quarter for the given time.
// Returns the first day of the quarter at 00:00:00.000000000 in the same timezone.
// Q1: January 1, Q2: April 1, Q3: July 1, Q4: October 1
//
// Example:
//
//	t := time.Date(2026, 6, 15, 15, 30, 45, 0, time.UTC)
//	start := lxtime.StartOfQuarter(t)
//	// start: 2026-04-01 00:00:00 +0000 UTC (Q2)
func StartOfQuarter(t time.Time) time.Time {
	m := time.Month(((int(t.Month())-1)/3)*3 + 1)

	return time.Date(
		t.Year(),
		m,
		1,
		0, 0, 0, 0,
		t.Location(),
	)
}
