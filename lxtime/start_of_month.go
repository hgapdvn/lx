package lxtime

import "time"

// StartOfMonth returns the start of the month for the given time.
// Returns day 1 at 00:00:00.000000000 of the same month.
//
// Example:
//
//	t := time.Date(2026, 4, 15, 15, 30, 45, 0, time.UTC)
//	start := lxtime.StartOfMonth(t)
//	// start: 2026-04-01 00:00:00 +0000 UTC
func StartOfMonth(t time.Time) time.Time {
	// Extract year, month from the input time
	year, month, _ := t.Date()

	// Create a new time for the first day of the month at midnight
	// Preserves the original timezone
	return time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
}
