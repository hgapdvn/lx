package lxtime

import "time"

// IsPast returns true if the given time is in the past (before now).
//
// Example:
//
//	t := lxtime.Ago(5, time.Minute)
//	isPast := lxtime.IsPast(t)
//	// isPast: true
//
//	future := lxtime.FromNow(5, time.Minute)
//	isFuturePast := lxtime.IsPast(future)
//	// isFuturePast: false
func IsPast(t time.Time) bool {
	return t.Before(time.Now())
}
