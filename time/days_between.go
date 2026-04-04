package lxtime

import "time"

// DaysBetween returns the number of calendar days between two times.
// The count includes both the start and end dates (exclusive of endpoints in terms of complete days).
// Negative values indicate t2 is before t1.
// Timezones are taken into account - the calculation is based on the date in each time's timezone.
//
// Example:
//
//	t1 := time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC)
//	t2 := time.Date(2026, 4, 6, 20, 45, 0, 0, time.UTC)
//	days := lxtime.DaysBetween(t1, t2)
//	// days: 2 (April 4 to April 6 is 2 days difference)
//
//	t3 := time.Date(2026, 4, 4, 0, 0, 0, 0, time.UTC)
//	t4 := time.Date(2026, 4, 4, 23, 59, 59, 0, time.UTC)
//	days := lxtime.DaysBetween(t3, t4)
//	// days: 0 (same day)
func DaysBetween(t1, t2 time.Time) int {
	loc := t1.Location()

	y1, m1, d1 := t1.In(loc).Date()
	y2, m2, d2 := t2.In(loc).Date()

	start := time.Date(y1, m1, d1, 0, 0, 0, 0, loc)
	end := time.Date(y2, m2, d2, 0, 0, 0, 0, loc)

	return int(end.Sub(start).Hours() / 24)
}
