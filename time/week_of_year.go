package lxtime

import "time"

// WeekOfYear returns the week number of the year for the given time.
// The week is defined as Monday through Sunday.
// Returns a value from 1 to 53.
// The first week is the one containing the first Monday of the year.
// Days before the first Monday are considered part of week 0 (but we return 1 for consistency).
//
// Example:
//
//	t := time.Date(2026, 1, 10, 10, 30, 0, 0, time.UTC) // Saturday in week 2
//	week := lxtime.WeekOfYear(t)
//	// week: 2
//
//	t2 := time.Date(2026, 1, 5, 10, 30, 0, 0, time.UTC) // Monday, first week
//	week2 := lxtime.WeekOfYear(t2)
//	// week2: 1
func WeekOfYear(t time.Time) int {
	// Get the start of the year
	year := t.Year()
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, t.Location())

	// Get the start of the current week
	weekStart := StartOfWeek(t)

	// Find the first Monday of the year
	// Start from Jan 1 and find the first Monday
	firstMonday := yearStart
	for firstMonday.Weekday() != time.Monday {
		firstMonday = firstMonday.AddDate(0, 0, 1)
	}

	// For dates before the first Monday, return 1
	if t.Before(firstMonday) {
		return 1
	}

	// Calculate days difference between week start and first Monday
	// Add 1 because we count weeks starting from 1
	daysDiff := int(weekStart.Sub(firstMonday).Hours() / 24)
	weekNumber := daysDiff/7 + 1

	return weekNumber
}
