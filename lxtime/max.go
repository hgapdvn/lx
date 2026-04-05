package lxtime

import "time"

// Max returns the later of two times.
// If both times are equal, either may be returned.
//
// Example:
//
//	t1 := time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC)
//	t2 := time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC)
//	later := lxtime.Max(t1, t2)
//	// later: 2026-04-15 14:00:00 +0000 UTC
func Max(a, b time.Time) time.Time {
	if a.After(b) {
		return a
	}
	return b
}
