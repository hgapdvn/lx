package lxtime

import "time"

// StartOfDay returns the start of the day for the given time.
// Returns the time at 00:00:00.000000000 on the same date.
//
// Example:
//
//	t := time.Date(2026, 4, 4, 15, 30, 45, 123456789, time.UTC)
//	start := lxtime.StartOfDay(t)
//	// start: 2026-04-04 00:00:00 +0000 UTC
func StartOfDay(t time.Time) time.Time {
	// Extract date components in the time's local timezone
	year, month, day := t.Date()

	// Reconstruct time at midnight, preserving timezone
	// This correctly handles non-UTC timezones, unlike Truncate() which uses UTC epoch
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}
