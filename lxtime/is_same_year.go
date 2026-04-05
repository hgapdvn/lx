package lxtime

import "time"

// IsSameYear compares two times and returns true if they are in the same year.
// The month, day, and time components are ignored.
// Timezones are taken into account - times in different timezones may be in different years.
//
// Example:
//
//	t1 := time.Date(2026, 1, 1, 10, 30, 0, 0, time.UTC)
//	t2 := time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC)
//	result := lxtime.IsSameYear(t1, t2)
//	// result: true (both in 2026)
//
//	t3 := time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC)
//	result := lxtime.IsSameYear(t1, t3)
//	// result: false (different years)
func IsSameYear(t1, t2 time.Time) bool {
	// Extract year components in their respective timezones
	y1 := t1.Year()
	y2 := t2.Year()

	// Compare years
	return y1 == y2
}
