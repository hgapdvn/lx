package lxtime

import "time"

// Clamp constrains a time to a [start, end] range.
// Returns start if t is before start, end if t is after end, or t if it's within the range.
// If start and end are equal, returns start.
// If start is after end, the behavior is undefined (start should be <= end).
//
// Example:
//
//	start := time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC)
//	end := time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC)
//	mid := time.Date(2026, 4, 15, 12, 0, 0, 0, time.UTC)
//
//	clamped := lxtime.Clamp(mid, start, end)
//	// clamped: 2026-04-15 12:00:00 +0000 UTC (unchanged, within range)
//
//	tooEarly := time.Date(2026, 4, 15, 8, 0, 0, 0, time.UTC)
//	clamped2 := lxtime.Clamp(tooEarly, start, end)
//	// clamped2: 2026-04-15 10:00:00 +0000 UTC (clamped to start)
//
//	tooLate := time.Date(2026, 4, 15, 16, 0, 0, 0, time.UTC)
//	clamped3 := lxtime.Clamp(tooLate, start, end)
//	// clamped3: 2026-04-15 14:00:00 +0000 UTC (clamped to end)
func Clamp(t, start, end time.Time) time.Time {
	if t.Before(start) {
		return start
	}
	if t.After(end) {
		return end
	}
	return t
}
