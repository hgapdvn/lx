package lxtime

import "time"

// IsSameWeek compares two times and returns true if they fall in the same Monday-Sunday week.
// The week boundaries are Monday at 00:00:00 to Sunday at 23:59:59.
// Timezones are taken into account - times in different timezones may be in different weeks.
//
// Example:
//
//	t1 := time.Date(2026, 4, 6, 10, 30, 0, 0, time.UTC) // Monday
//	t2 := time.Date(2026, 4, 8, 15, 45, 0, 0, time.UTC) // Wednesday
//	result := lxtime.IsSameWeek(t1, t2)
//	// result: true (same week: Mon Apr 6 - Sun Apr 12)
//
//	t3 := time.Date(2026, 4, 13, 10, 0, 0, 0, time.UTC) // Monday of next week
//	result := lxtime.IsSameWeek(t1, t3)
//	// result: false (different weeks)
func IsSameWeek(t1, t2 time.Time) bool {
	// Get the start of the week for both times
	start1 := StartOfWeek(t1)
	start2 := StartOfWeek(t2)

	// Compare the year, month, and day of the week starts
	y1, m1, d1 := start1.Date()
	y2, m2, d2 := start2.Date()

	return y1 == y2 && m1 == m2 && d1 == d2
}
