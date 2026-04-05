package lxtime

import "time"

// IsToday returns true if the given time is today (in the local timezone).
// This compares the local calendar dates, so times that are on the same day in
// the same location are considered "today" even if they're different UTC dates.
//
// Example:
//
//	t := time.Now()
//	if lxtime.IsToday(t) {
//		// t is today
//	}
func IsToday(t time.Time) bool {
	now := time.Now()
	// Compare local calendar dates
	y1, m1, d1 := t.Date()
	y2, m2, d2 := now.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
