package lxtime

import "time"

// IsSameMonth compares two times and returns true if they are in the same year and month.
// The day and time components are ignored.
// Timezones are taken into account - times in different timezones may be in different months.
//
// Example:
//
//	t1 := time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC)
//	t2 := time.Date(2026, 4, 30, 23, 59, 59, 0, time.UTC)
//	result := lxtime.IsSameMonth(t1, t2)
//	// result: true (both in April 2026)
//
//	t3 := time.Date(2026, 5, 1, 0, 0, 0, 0, time.UTC)
//	result := lxtime.IsSameMonth(t1, t3)
//	// result: false (different months)
func IsSameMonth(t1, t2 time.Time) bool {
	// Extract year and month components in their respective timezones
	y1, m1, _ := t1.Date()
	y2, m2, _ := t2.Date()

	// Compare year and month
	return y1 == y2 && m1 == m2
}
