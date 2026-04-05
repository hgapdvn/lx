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
	// Get the current day of week (0=Sunday, 1=Monday, ..., 6=Saturday)
	currentWeekday := t.Weekday()

	// Calculate how many days back to Monday
	// Monday is always 1, so: Monday(1) - Monday(1) = 0, Tuesday(2) - Monday(1) = 1, etc.
	daysBackToMonday := int(currentWeekday - time.Monday)

	// Special case: Sunday (0) gives us 0 - 1 = -1
	// For Sunday, we want to go back 6 days to the previous week's Monday
	if daysBackToMonday < 0 {
		daysBackToMonday = 6
	}

	// Extract date components in the time's local timezone
	year, month, day := t.Date()

	// Create start of current day, then subtract days to get Monday
	// Using Date() correctly handles non-UTC timezones
	startOfDay := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	return startOfDay.AddDate(0, 0, -daysBackToMonday)
}
