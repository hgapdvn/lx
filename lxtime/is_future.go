package lxtime

import "time"

// IsFuture returns true if the given time is in the future (after now).
//
// Example:
//
//	t := lxtime.FromNow(5, time.Minute)
//	isFuture := lxtime.IsFuture(t)
//	// isFuture: true
//
//	past := lxtime.Ago(5, time.Minute)
//	isPastFuture := lxtime.IsFuture(past)
//	// isPastFuture: false
func IsFuture(t time.Time) bool {
	return t.After(time.Now())
}
