package lxtime

import "time"

// IsTomorrow returns true if the given time is tomorrow (in the local timezone).
// This compares the local calendar dates, so times that are on the same day in
// the same location are considered "tomorrow" even if they're different UTC dates.
//
// Example:
//
//	t := time.Now().AddDate(0, 0, 1)
//	if lxtime.IsTomorrow(t) {
//		// t is tomorrow
//	}
func IsTomorrow(t time.Time) bool {
	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)
	// Compare local calendar dates
	y1, m1, d1 := t.Date()
	y2, m2, d2 := tomorrow.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
