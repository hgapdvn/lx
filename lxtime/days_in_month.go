package lxtime

import "time"

// TotalDaysInMonth returns the number of days in the month of the given time.
// Correctly handles leap years for February.
//
// Example:
//
//	t := time.Date(2026, 4, 15, 0, 0, 0, 0, time.UTC)
//	days := lxtime.TotalDaysInMonth(t)
//	// days: 30
//
//	leapYearTime := time.Date(2024, 2, 15, 0, 0, 0, 0, time.UTC)
//	leapDays := lxtime.TotalDaysInMonth(leapYearTime)
//	// leapDays: 29
func TotalDaysInMonth(t time.Time) int {
	year, month, _ := t.Date()

	// Get the first day of next month, then subtract 1 day
	firstDayOfNextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, t.Location())
	lastDayOfCurrentMonth := firstDayOfNextMonth.Add(-time.Nanosecond)

	// Return the day of the last day of the current month
	return lastDayOfCurrentMonth.Day()
}
