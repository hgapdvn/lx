package lxtime

import "time"

// StartOfYear returns the start of the year for the given time.
// Returns January 1 at 00:00:00.000000000 of the same year.
//
// Example:
//
//	t := time.Date(2026, 6, 15, 15, 30, 45, 0, time.UTC)
//	start := lxtime.StartOfYear(t)
//	// start: 2026-01-01 00:00:00 +0000 UTC
func StartOfYear(t time.Time) time.Time {
	// Extract year from the input time
	year := t.Year()

	// Create a new time for January 1 at midnight, preserving timezone
	return time.Date(year, time.January, 1, 0, 0, 0, 0, t.Location())
}
