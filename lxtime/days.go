package lxtime

import "time"

// Days returns a duration representing n days.
//
// Example:
//
//	duration := lxtime.Days(5)
//	// duration: 120h0m0s (5 days)
func Days(n int) time.Duration {
	return time.Duration(n) * 24 * time.Hour
}
