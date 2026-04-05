package lxtime

import "time"

// IsWeekend returns true if the given time falls on a weekend (Saturday or Sunday).
// Returns false for weekdays (Monday-Friday).
//
// Example:
//
//	t := time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC) // Saturday
//	if lxtime.IsWeekend(t) {
//		// t is a weekend day
//	}
func IsWeekend(t time.Time) bool {
	day := t.Weekday()
	return day == time.Saturday || day == time.Sunday
}
