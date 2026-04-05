package lxtime

import "time"

// EndOfDay returns the end of the day for the given time.
// Returns the time at 23:59:59.999999999 on the same date.
//
// Example:
//
//	t := time.Date(2026, 4, 4, 15, 30, 45, 123456789, time.UTC)
//	end := lxtime.EndOfDay(t)
//	// end: 2026-04-04 23:59:59.999999999 +0000 UTC
func EndOfDay(t time.Time) time.Time {
	// Extract date components in the time's local timezone
	year, month, day := t.Date()

	// Reconstruct time at end of day (23:59:59.999999999), preserving timezone
	// This correctly handles non-UTC timezones, unlike Truncate() which uses UTC epoch
	return time.Date(year, month, day, 23, 59, 59, 999999999, t.Location())
}
