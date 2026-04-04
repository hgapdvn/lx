package lxtime

import "time"

// IsWeekDay returns true if the given time falls on a weekday (Monday-Friday).
// Returns false for Saturday and Sunday.
//
// Example:
//
//	t := time.Date(2026, 4, 6, 12, 0, 0, 0, time.UTC) // Monday
//	if lxtime.IsWeekDay(t) {
//		// t is a weekday
//	}
func IsWeekDay(t time.Time) bool {
	day := t.Weekday()
	return day != time.Saturday && day != time.Sunday
}
