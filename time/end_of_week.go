package lxtime

import "time"

// EndOfWeek returns the end of the week for the given time.
// Returns Sunday at 23:59:59.999999999 of the same week.
// If the given time is a Sunday, returns the end of that Sunday.
//
// Example:
//
//	t := time.Date(2026, 4, 8, 15, 30, 45, 0, time.UTC) // Wednesday
//	end := lxtime.EndOfWeek(t)
//	// end: 2026-04-12 23:59:59.999999999 +0000 UTC (Sunday)
func EndOfWeek(t time.Time) time.Time {
	weekday := t.Weekday()
	daysToSunday := (7 - int(weekday)) % 7
	return t.Truncate(24*time.Hour).AddDate(0, 0, daysToSunday).Add(24*time.Hour - 1*time.Nanosecond)
}
