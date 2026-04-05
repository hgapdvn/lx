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
	// Get the current day of week (0=Sunday, 1=Monday, ..., 6=Saturday)
	currentWeekday := t.Weekday()

	// Calculate how many days forward to Sunday
	// Sunday is 0, so: (7 - 0) % 7 = 0 (already Sunday), (7 - 1) % 7 = 6 (Monday needs 6 days), etc.
	daysForwardToSunday := (7 - int(currentWeekday)) % 7

	// Extract date components in the time's local timezone
	year, month, day := t.Date()

	// Create start of current day, then add days to get Sunday
	// Using Date() correctly handles non-UTC timezones
	startOfDay := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	sundayStart := startOfDay.AddDate(0, 0, daysForwardToSunday)

	// Subtract 1 nanosecond to get 23:59:59.999999999 instead of 00:00:00
	return sundayStart.Add(24*time.Hour - 1*time.Nanosecond)
}
