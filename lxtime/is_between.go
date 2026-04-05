package lxtime

import "time"

// IsBetween returns true if time is between start and end, inclusive.
// If start and end are equal, it returns true.
//
// Example:
//
//	t := time.Now()
//	if lxtime.IsBetween(t, time.Now().Add(-1*time.Hour), time.Now().Add(1*time.Hour)) {
//		// t is between now-1h and now+1h
//	}
func IsBetween(t, start, end time.Time) bool {
	return (t.Equal(start) || t.After(start)) && (t.Equal(end) || t.Before(end))
}
