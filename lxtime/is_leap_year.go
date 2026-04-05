package lxtime

import "time"

// IsLeapYear returns true if the year in the given time is a leap year.
// A year is a leap year if:
//   - It is divisible by 4 AND
//   - (It is not divisible by 100 OR it is divisible by 400)
//
// This means:
//   - 2024 is a leap year (divisible by 4, not by 100)
//   - 2000 is a leap year (divisible by 400)
//   - 1900 is not a leap year (divisible by 100 but not by 400)
//   - 2001 is not a leap year (not divisible by 4)
//
// Example:
//
//	t := time.Date(2024, 2, 15, 0, 0, 0, 0, time.UTC)
//	isLeap := lxtime.IsLeapYear(t)
//	// isLeap: true
//
//	t2 := time.Date(2026, 2, 15, 0, 0, 0, 0, time.UTC)
//	isLeap2 := lxtime.IsLeapYear(t2)
//	// isLeap2: false
func IsLeapYear(t time.Time) bool {
	year := t.Year()

	// A year is a leap year if:
	// - divisible by 400, OR
	// - divisible by 4 AND not divisible by 100
	return (year%400 == 0) || (year%4 == 0 && year%100 != 0)
}
