package lxtime

import "time"

// IsYesterday returns true if the given time is yesterday (in the local timezone).
// This compares the local calendar dates, so times that are on the same day in
// the same location are considered "yesterday" even if they're different UTC dates.
//
// Example:
//
//	t := time.Now().AddDate(0, 0, -1)
//	if lxtime.IsYesterday(t) {
//		// t is yesterday
//	}
func IsYesterday(t time.Time) bool {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	// Compare local calendar dates
	y1, m1, d1 := t.Date()
	y2, m2, d2 := yesterday.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
