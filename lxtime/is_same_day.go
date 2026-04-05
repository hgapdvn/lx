package lxtime

import "time"

// IsSameDay compares two times by date and returns true if they are on the same day.
// The time components (hours, minutes, seconds, nanoseconds) are ignored.
// Timezones are taken into account - times in different timezones may be different days.
//
// Example:
//
//	t1 := time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC)
//	t2 := time.Date(2026, 4, 4, 23, 59, 59, 0, time.UTC)
//	result := lxtime.IsSameDay(t1, t2)
//	// result: true (same day)
//
//	t3 := time.Date(2026, 4, 5, 10, 30, 0, 0, time.UTC)
//	result := lxtime.IsSameDay(t1, t3)
//	// result: false (different days)
func IsSameDay(t1, t2 time.Time) bool {
	// Extract date components in their respective timezones
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()

	// Compare year, month, and day
	return y1 == y2 && m1 == m2 && d1 == d2
}
