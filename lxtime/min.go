package lxtime

import "time"

// Min returns the earlier of two times.
// If both times are equal, either may be returned.
//
// Example:
//
//	t1 := time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC)
//	t2 := time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC)
//	earlier := lxtime.Min(t1, t2)
//	// earlier: 2026-04-15 10:00:00 +0000 UTC
func Min(a, b time.Time) time.Time {
	if a.Before(b) {
		return a
	}
	return b
}
