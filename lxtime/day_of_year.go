package lxtime

import "time"

// DayOfYear returns the day number of the year for the given time.
// Returns a value from 1 to 366, where 1 is January 1st and 366 is December 31st (in leap years).
//
// Example:
//
//	t := time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC)
//	day := lxtime.DayOfYear(t)
//	// day: 94 (April 4 is the 94th day of 2026)
//
//	t2 := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
//	day2 := lxtime.DayOfYear(t2)
//	// day2: 1 (January 1st is the first day)
func DayOfYear(t time.Time) int {
	// time.Time has a YearDay() method that returns the day of year from 1-366
	return t.YearDay()
}
