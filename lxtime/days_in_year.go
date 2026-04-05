package lxtime

import "time"

// TotalDaysInYear returns the number of days in the year of the given time.
// Returns 366 for leap years, 365 for non-leap years.
//
// Example:
//
//	t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
//	days := lxtime.TotalDaysInYear(t)
//	// days: 366 (2024 is a leap year)
//
//	t2 := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
//	days2 := lxtime.TotalDaysInYear(t2)
//	// days2: 365 (2026 is not a leap year)
func TotalDaysInYear(t time.Time) int {
	if IsLeapYear(t) {
		return 366
	}
	return 365
}
