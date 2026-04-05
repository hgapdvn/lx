package lxtime

import "time"

// EndOfYear returns December 31 at 23:59:59.999999999 of the same year.
func EndOfYear(t time.Time) time.Time {
	year := t.Year()
	return time.Date(year, time.December, 31, 23, 59, 59, 999999999, t.Location())
}
