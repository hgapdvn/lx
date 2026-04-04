package lxtime

import "time"

// Weeks returns a duration representing n weeks.
//
// Example:
//
//	duration := lxtime.Weeks(2)
//	// duration: 336h0m0s (2 weeks)
func Weeks(n int) time.Duration {
	return time.Duration(n) * 7 * 24 * time.Hour
}
