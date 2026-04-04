package lxtime

import "time"

// StartOfWeek returns the start of the week for the given time.
// Returns Monday at 00:00:00.000000000 of the same week.
// If the given time is a Monday, returns the start of that Monday.
//
// Example:
//
//	t := time.Date(2026, 4, 8, 15, 30, 45, 0, time.UTC) // Wednesday
//	start := lxtime.StartOfWeek(t)
//	// start: 2026-04-06 00:00:00 +0000 UTC (Monday)
func StartOfWeek(t time.Time) time.Time {
	weekday := t.Weekday()
	daysToMonday := int(weekday - time.Monday)
	if daysToMonday < 0 {
		daysToMonday = 6
	}
	return t.Truncate(24*time.Hour).AddDate(0, 0, -daysToMonday)
}
