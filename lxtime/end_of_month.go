package lxtime

import "time"

// EndOfMonth returns the end of the month for the given time.
// Returns the last day of the month at 23:59:59.999999999.
//
// Example:
//
//	t := time.Date(2026, 4, 15, 15, 30, 45, 0, time.UTC)
//	end := lxtime.EndOfMonth(t)
//	// end: 2026-04-30 23:59:59.999999999 +0000 UTC
func EndOfMonth(t time.Time) time.Time {
	// Extract year and month from the input time
	year, month, _ := t.Date()

	// Create time for the first day of NEXT month at midnight
	// Example: April 15 -> May 1, 00:00:00
	firstDayOfNextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, t.Location())

	// Subtract 1 day to get the last day of current month at midnight
	// Example: May 1 - 1 day = April 30, 00:00:00
	lastDayOfMonth := firstDayOfNextMonth.AddDate(0, 0, -1)

	// Subtract 1 nanosecond to get 23:59:59.999999999 instead of 00:00:00
	// Example: April 30, 00:00:00 - 1ns = April 29, 23:59:59.999999999
	return lastDayOfMonth.Add(24*time.Hour - 1*time.Nanosecond)
}
